package mqtt

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/mahdi-cpp/api-go-emqx/model"
	"github.com/mahdi-cpp/api-go-emqx/repository"
	"strings"
)

var client MQTT.Client

// Handler برای دریافت پیام‌های دریافتی
var messageHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

	//fmt.Printf("Received message on topic %s: %s\n", msg.Topic(), msg.Payload())

	if strings.Compare(msg.Topic(), "error") == 0 {
		//fmt.Printf("error topic\n")
	} else {

		// Create an instance of Packet
		var packet model.Packet

		// Unmarshal the JSON data into the Packet struct
		err := json.Unmarshal(msg.Payload(), &packet)
		if err != nil {
			//log.Fatalf("Error unmarshaling JSON: %v", err)
		} else {
			//3. Print each field individually
			fmt.Println("-----------------------------------------")
			fmt.Printf("ID: %s\n", packet.Id)
			//fmt.Printf("Message: %s\n", packet.Message)
			//fmt.Printf("Temperature: %s\n", packet.Temperature)
			//fmt.Printf("Value: %d\n", packet.Value)
		}
		repository.Add(msg.Topic(), string(msg.Payload()))
		repository.AddTemperature(packet.Id, packet.Value)
	}
}

func Start() MQTT.Client {

	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://192.168.1.114:1883")    // آدرس بروکر عمومی تست
	opts.SetClientID("go-mqtt-client")            // شناسه کلاینت
	opts.SetDefaultPublishHandler(messageHandler) // هندلر پیش‌فرض پیام‌ها
	opts.SetUsername("admin")
	opts.SetPassword("admin@123456")

	client = MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Connected to MQTT broker!")

	if token := client.Subscribe("temperature/topic", 1, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("error", 1, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Println("Subscribed to topic: temperature/topic")

	// ارسال پیام به Topic
	token := client.Publish("inTopic", 1, false, "light off")
	token.Wait()
	fmt.Println("Message published!")

	// منتظر سیگنال‌های خروج باش (Ctrl+C)
	//sigChan := make(chan os.Signal, 1)
	//signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	//<-sigChan

	// قطع اتصال
	client.Disconnect(250)
	fmt.Println("Disconnected from MQTT broker.")

	return client

}

func SendMQTT(msg string) {
	token := client.Publish("inTopic", 1, false, msg)
	token.Wait()
	fmt.Println("Message published!")
}
