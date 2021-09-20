package models

import (
	"time"
)

type Item struct {
	Item_id     uint   `gorm:"primaryKey" json:"line_item_id"`
	Item_code   string `json:"item_code"`
	Description string
	Quantity    int
	Order_id    uint `json:"order_id"`
}

type Order struct {
	Order_id      uint      `gorm:"primaryKey" json:"order_id"`
	Customer_name string    `json:"customer_name"`
	Ordered_at    time.Time `json:"ordered_at"`
	Items         []Item    `gorm:"foreignKey:order_id;references:order_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
