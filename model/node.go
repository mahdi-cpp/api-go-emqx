package model

import "time"

type Node struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	Topic     string `json:"topic"`
	MessageID uint16 `json:"message_id"`
	Payload   string `json:"payload"`
	Value     int32  `json:"value"`
}
