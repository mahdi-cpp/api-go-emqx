package model

type Lamp struct {
	ID    uint   `gorm:"primarykey"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Value int32  `json:"value"`
}
