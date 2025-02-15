package websocket

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/mahdi-cpp/api-go-emqx/model"
	"github.com/mahdi-cpp/api-go-emqx/repository"
	"github.com/mahdi-cpp/api-go-emqx/utils"
	"net/http"
	"strings"
	"sync"
)

var mqttClient MQTT.Client

var upgrader = websocket.Upgrader{}

var clients = make(map[string]*websocket.Conn) // Store connections with IDs
var mu sync.Mutex
var conn *websocket.Conn // Mutex for safe concurrent access

func AddMQTT(client MQTT.Client) {
	mqttClient = client
}

func handleConnection(w http.ResponseWriter, r *http.Request) {

	var err error
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()

	// Generate a unique client ID
	clientID := uuid.New().String() // Generate a new UUID

	// Store the connection with the client ID
	mu.Lock()
	clients[clientID] = conn
	mu.Unlock()

	fmt.Printf("Client connected: %s\n", clientID)

	for {
		// Read message from the client
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error while reading message:", err)
			break
		}

		fmt.Printf("Received message from messageType:%d:  clientID:%s: %s\n", messageType, clientID, msg)

		var text = string(msg)
		if strings.Contains(text, "light on") {
			sendMQTT(text)
			sendText("light is on")
		} else if strings.Contains(text, "light off") {
			sendMQTT(text)
			sendText("light is off")
		} else if strings.Contains(text, "lamps") {
			err := processLamps()
			if err != nil {
				return
			}
		} else if strings.Contains(text, "tmp") {
			processTemperature()
		} else {
			sendText("command is wrong!")
		}
	}

	// Remove the client from the map when they disconnect
	mu.Lock()
	delete(clients, clientID)
	mu.Unlock()
	fmt.Printf("Client disconnected: %s\n", clientID)
}

func Start() {

	http.HandleFunc("/ws", handleConnection)

	// Start the WebSocket server in a separate goroutine
	go func() {
		fmt.Println("WebSocket server started on :8097")
		if err := http.ListenAndServe(":8097", nil); err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()
}

func sendMQTT(msg string) {
	token := mqttClient.Publish("inTopic", 1, false, msg)
	token.Wait()
	fmt.Println("Message published!")
}

func sendText(text string) {

	var object model.Object
	object.Type = "text"
	object.JsonString = text

	jsonBytes, err := utils.ConvertObjectToBytes(object)
	if err != nil {
		fmt.Println("Error converting object to bytes:", err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
		fmt.Println("Error while writing message:", err)
	}
}

func processLamps() error {

	var lamps []model.Lamp

	var a = model.Lamp{Name: "Living Room", Icon: "lap-4-100.png", Value: 1}
	var b = model.Lamp{Name: "Ali Room", Icon: "icons8-lights-100.png", Value: 1}
	var c = model.Lamp{Name: "Mahdi Room", Icon: "icons8-table-lights-100.png", Value: 0}
	var d = model.Lamp{Name: "Kitchen 3", Icon: "lap-4-100.png", Value: 0}
	var a1 = model.Lamp{Name: "Room Lamp", Icon: "lap-4-100.png", Value: 1}
	var b1 = model.Lamp{Name: "Lamp 21", Icon: "lap-4-100.png", Value: 1}
	var c1 = model.Lamp{Name: "Lamp 31", Icon: "lap-4-100.png", Value: 0}
	var d1 = model.Lamp{Name: "Lamp 13", Icon: "lap-4-100.png", Value: 1}

	lamps = append(lamps, a)
	lamps = append(lamps, b)
	lamps = append(lamps, c)
	lamps = append(lamps, d)
	lamps = append(lamps, a1)
	lamps = append(lamps, b1)
	lamps = append(lamps, c1)
	lamps = append(lamps, d1)

	// Convert the slice of Temperature structs to JSON
	jsonData, err := json.Marshal(lamps)
	if err != nil {
		return err
	}

	var object model.Object
	object.Type = "lamps"
	object.JsonString = "{lamps:" + string(jsonData) + "}"

	// Convert the Object to a byte slice
	jsonBytes, err := utils.ConvertObjectToBytes(object)
	if err != nil {
		fmt.Println("Error converting object to bytes:", err)
		return err
	}

	if err := conn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
		fmt.Println("Error while writing message:", err)
	}

	return nil
}

func processTemperature() {

	//Read temperatures from the last hour and get the JSON string
	jsonResult, err := repository.ReadTemperatures()
	if err != nil {
		fmt.Println("Error reading recent temperatures:", err)
		return
	}

	var object model.Object
	object.Type = "temperature"
	object.JsonString = "{temperatures:" + jsonResult + "}"

	// Convert the Object to a byte slice
	jsonBytes, err := utils.ConvertObjectToBytes(object)
	if err != nil {
		fmt.Println("Error converting object to bytes:", err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
		fmt.Println("Error while writing message:", err)
	}
}
