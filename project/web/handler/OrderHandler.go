package handler

import (
	"net/http"
	"project/web"
)

type OrderHandler struct {
	OrderRepo   *OrderRepo
	ProductRepo *ProductRepo
}

func (orderHandler *OrderHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	order, orderItems, err := orderHandler.OrderRepo.PlaceOrder(userId)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	orderDto := web.ConvertToDto(order, OrderDto{})
	var orderItemsDto []OrderItemDto
	for _, orderItem := range orderItems {
		orderItemDto:= web.ConvertToDto(orderItem, OrderItemDto{})
		web.AssignProductNameToOrder(&orderItemDto, orderHandler.ProductRepo, orderItem.Id)
		orderItemsDto=append(orderItemsDto, orderItemDto)
	}
	orderDto.OrderItems=orderItemsDto
	web.CreateResponse(w,nil,orderDto)
}
