package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/mahdi-cpp/api-go-emqx/cache"
	"github.com/mahdi-cpp/api-go-emqx/config"
	"github.com/mahdi-cpp/api-go-emqx/repository"
)

// Handler برای دریافت پیام‌های دریافتی
var messageHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message on topic %s: %s\n", msg.Topic(), msg.Payload())
}

func main() {

	config.LayoutInit()
	repository.InitModels()
	cache.ReadIcons()

	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://192.168.1.114:1883")    // آدرس بروکر عمومی تست
	opts.SetClientID("go-mqtt-client")            // شناسه کلاینت
	opts.SetDefaultPublishHandler(messageHandler) // هندلر پیش‌فرض پیام‌ها
	opts.SetUsername("admin")
	opts.SetPassword("admin@123456")

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Connected to MQTT broker!")

	if token := client.Subscribe("temperature/topic", 1, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Subscribed to topic: temperature/topic")

	// ارسال پیام به Topic
	//token := client.Publish("temperature/topic", 1, false, "Hello from Go!")
	//token.Wait()
	//fmt.Println("Message published!")

	// منتظر سیگنال‌های خروج باش (Ctrl+C)
	//sigChan := make(chan os.Signal, 1)
	//signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	//<-sigChan

	Run()

	// قطع اتصال
	client.Disconnect(250)
	fmt.Println("Disconnected from MQTT broker.")

}
