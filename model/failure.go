package model

import "time"

type Failure struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	Topic     string `json:"topic"`
	NodeID    string `json:"NodeID"`
	Message   string `json:"message"`
}
