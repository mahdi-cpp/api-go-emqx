package repository

import (
	"github.com/mahdi-cpp/api-go-emqx/config"
	"github.com/mahdi-cpp/api-go-emqx/model"
)

func InitNodes() {
	//config.DB.Create(&model.Node{Name: "Esp8266", Value: 12})
	//config.DB.Create(&model.Node{Name: "Esp8266", Value: 12})
	//config.DB.Create(&model.Node{Name: "Esp8266", Value: 12})
	//config.DB.Create(&model.Node{Name: "Esp8266", Value: 12})
}

func Add(topic string, payload string) {
	config.DB.Create(&model.Node{Topic: topic, Payload: payload})
}

func AddTemperature(NodeId string, value int32) {
	config.DB.Create(&model.Temperature{NodeId: NodeId, Value: value})
}
