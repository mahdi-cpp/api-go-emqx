package main

import (
	"github.com/mahdi-cpp/api-go-emqx/cache"
	"github.com/mahdi-cpp/api-go-emqx/config"
	"github.com/mahdi-cpp/api-go-emqx/mqtt"
	"github.com/mahdi-cpp/api-go-emqx/repository"
	"github.com/mahdi-cpp/api-go-emqx/websocket"
)

func main() {

	config.LayoutInit()
	repository.InitModels()
	cache.ReadIcons()

	config.DatabaseInit()
	//repository.InitNodes()

	mqttClient := mqtt.Start()

	websocket.AddMQTT(mqttClient)
	websocket.Start()

	Run()

}

func SendMqtt() {

}
