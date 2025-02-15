package model

import "time"

type Temperature struct {
	ID        uint      `gorm:"primarykey"`
	NodeId    string    `json:"nodeId"`
	Value     int32     `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
}
