package model

import (
	"time"
)

type Product struct {
	Id          uint64 `gorm:"primary_key"`
	Name        string
	Brand       string
	Price       float64 `gorm:"type:decimal(10,2);default:0.00"`
	Inventory   uint
	Description string
	CategoryId  uint64
}

type ProductDto struct {
	Id           uint64  `json:"id"`
	Name         string  `json:"name"`
	Brand        string  `json:"brand"`
	Price        float64 `json:"price"`
	Inventory    int     `json:"inventory"`
	Description  string  `json:"description"`
	CategoryName string  `json:"category_name"`
}
type Category struct {
	Id   uint64 `gorm:"primary_key"`
	Name string
}

type Cart struct {
	Id          uint64  `gorm:"primary_key"`
	TotalAmount float64 `gorm:"type:decimal(10,2);default:0.00"`
	UserId      string
}

type CartDto struct {
	Id          uint64        `json:"id"`
	TotalAmount float64       `json:"total_amount"`
	UserId      string        `json:"user_id"`
	CartItems   []CartItemDto `json:"cart_items"`
}
type CartItem struct {
	Id         uint64 `gorm:"primary_key"`
	Quantity   uint
	UnitPrice  float64 `gorm:"type:decimal(10,2);default:0.00"`
	TotalPrice float64 `gorm:"type.decimal(10,2);default:0.00"`
	ProductId  uint64
	CartId     uint64
}

type CartItemDto struct {
	Id          uint64  `json:"id"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
	ProductName string  `json:"product_name"`
}

type Order struct {
	Id          uint64 `gorm:"primary_key"`
	OrderDate   time.Time
	TotalPrice  float64 `gorm:"type.decimal(10,2);default:0.00"`
	OrderStatus string
	UserId      string
}
type OrderDto struct {
	Id          uint64         `json:"id"`
	OrderDate   time.Time      `json:"order_date"`
	TotalPrice  float64        `json:"total_price"`
	OrderStatus string         `json:"order_status"`
	OrderItems  []OrderItemDto `json:"order_items"`
}
type OrderItem struct {
	Id         uint64 `gorm:"primary_key"`
	OrderId    uint64
	ProductId  uint64
	Quantity   uint
	UnitPrice  float64 `gorm:"type:decimal(10,2);default:0.00"`
	TotalPrice float64 `gorm:"type.decimal(10,2);default:0.00"`
}
type OrderItemDto struct {
	Id          uint64  `json:"id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}

type Response struct {
	Status    string      `json:"status"`
	TimeStamp time.Time   `json:"timeStamp"`
	Data      interface{} `json:"data"`
	Error     interface{} `json:"error"`
}

type Request struct {
	SearchKey string `json:"search_key"`
}
