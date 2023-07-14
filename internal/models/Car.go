package models

type Car struct {
	ID    uint     `gorm:"primary key;autoIncrement" json:"id"`
	Name  *string  `json:"name"`
	Price *float64 `json:"price"`
}
