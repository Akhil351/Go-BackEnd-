package repo

import (
	"gorm.io/gorm"
)

type OrderItemRepo struct {
	Repo         *gorm.DB
	CartItemRepo *CartItemRepo
	ProductRepo  *ProductRepo
}

func (orderItemRepo *OrderItemRepo) CreateOrderItems(orderId uint64, cartId uint64) ([]OrderItem, error) {
	var orderItems []OrderItem
	cartItems, err := orderItemRepo.CartItemRepo.FindAllCartItemsByCartId(cartId)
	if err != nil {
		return orderItems, err
	}
	for _, cartItem := range cartItems {
		product, err := orderItemRepo.ProductRepo.FindProductById(cartItem.ProductId)
		if err != nil {
			return orderItems, err
		}
		product.Inventory = product.Inventory - cartItem.Quantity
		_, err = orderItemRepo.ProductRepo.AddProduct(product)
		if err != nil {
			return orderItems, err
		}
		orderItem := OrderItem{
			OrderId:    orderId,
			ProductId:  product.Id,
			UnitPrice:  cartItem.UnitPrice,
			TotalPrice: cartItem.TotalPrice,
			Quantity:   cartItem.Quantity,
		}
		orderItems = append(orderItems, orderItem)
	}
	if err := orderItemRepo.Repo.Create(&orderItems).Error; err != nil {
		return orderItems, err
	}
	orderItemRepo.CartItemRepo.DeleteAllCartItem(cartItems)
	return orderItems, nil
}

func (orderItemRepo *OrderItemRepo) GetOrderItemsByOrderId(orderId uint64) ([]OrderItem, error) {
	var orderItems []OrderItem
	if err := orderItemRepo.Repo.Where("order_id=?", orderId).Find(&orderItems).Error; err != nil {
		return orderItems, err
	}
	return orderItems, nil
}
