package model

import "time"

type Product struct {
	Id          uint `gorm:"primary_key"`
	Name        string
	Brand       string
	Price       float64
	Inventory   uint
	Description string
	CategoryId  uint
}
type Category struct {
	Id   uint `gorm:"primary_key"`
	Name string
}

type ProductDto struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Brand        string  `json:"brand"`
	Price        float64 `json:"price"`
	Inventory    int     `json:"inventory"`
	Description  string  `json:"description"`
	CategoryName string  `json:"category_name"`
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
