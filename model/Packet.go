package model

type Packet struct {
	Id          string `json:"id"`
	Message     string `json:"message"`
	Temperature string `json:"temperature"`
	Value       int32  `json:"value"`
}
