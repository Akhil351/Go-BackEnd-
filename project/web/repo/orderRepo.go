package repo

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type OrderRepo struct {
	Repo          *gorm.DB
	CartRepo      *CartRepo
	OrderItemRepo *OrderItemRepo
}

func (orderRepo *OrderRepo) PlaceOrder(userId string) (Order,[]OrderItem, error) {
	var order Order
	cart, err := orderRepo.CartRepo.FindCartByUserId(userId)
	if err != nil {
		return order,nil, err
	}
	if cart.TotalAmount == 0 {
		return order,nil, errors.New("The cart is empty. No items to order.")
	}
	order, err = orderRepo.createOrder(cart)
	if err != nil {
		return order, nil,err
	}
	orderItems, err := orderRepo.OrderItemRepo.CreateOrderItems(order.Id, cart.Id)
	if err != nil {
		return order,nil, err
	}
	cart.TotalAmount = 0
	err = orderRepo.CartRepo.SaveCart(cart)
	if err != nil {
		return order,nil, err
	}
	return order,orderItems, nil

}
func (orderRepo *OrderRepo) createOrder(cart Cart) (Order, error) {
	order := Order{
		OrderDate:   time.Now(),
		OrderStatus: "Delivered",
		UserId:      cart.UserId,
		TotalPrice:  cart.TotalAmount,
	}
	if err := orderRepo.Repo.Save(&order).Error; err != nil {
		return order, err
	}
	return order, nil


}
func (orderRepo *OrderRepo) FindOrderByUserId(userId string)([]Order,error){
	var order []Order
	if err:=orderRepo.Repo.Where("user_id=?",userId).Find(&order).Error; err!=nil{
		return order,err
	}
	return order,nil
}
