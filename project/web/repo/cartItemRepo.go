package repo

import (
	"gorm.io/gorm"
)

type CartItemRepo struct {
	Repo        *gorm.DB
	CartRepo    *CartRepo
	ProductRepo *ProductRepo
}

func (cartItemRepo *CartItemRepo) AddItemToCart(userId string, productId uint64, quantity int, methodName string) error {
	cart, err := cartItemRepo.CartRepo.FindCartByUserId(userId)
	if err != nil {
		return err
	}
	product, err := cartItemRepo.ProductRepo.FindProductById(productId)
	if err != nil {
		return err
	}
	cartItem, err := cartItemRepo.findCartByProductIdAndCartId(cart.Id, productId)
	if err != nil {
		cartItem.CartId = cart.Id
		cartItem.ProductId = productId
		cartItem.Quantity = uint(quantity)
		cartItem.UnitPrice = product.Price
	} else {
		if methodName == "Add" {
			cartItem.Quantity = cartItem.Quantity + uint(quantity)
		} else if methodName == "Update" {
			cartItem.Quantity = uint(quantity)
		}
	}
	cartItem.TotalPrice = cartItem.UnitPrice * float64(cartItem.Quantity)
	if err := cartItemRepo.Repo.Save(&cartItem).Error; err != nil {
		return err
	}
	cart.TotalAmount = cartItemRepo.updateTotalAmount(cart.Id)
	err = cartItemRepo.CartRepo.SaveCart(cart)
	if err != nil {
		return err
	}
	return nil

}
func (cartItemRepo *CartItemRepo) findCartByProductIdAndCartId(cartId uint64, productId uint64) (CartItem, error) {
	var cartItem CartItem
	if err := cartItemRepo.Repo.Where("product_id=? and cart_id=?", productId, cartId).First(&cartItem).Error; err != nil {
		return cartItem, err
	}
	return cartItem, nil
}
func (cartItemRepo *CartItemRepo) FindAllCartItemsByCartId(cartId uint64) ([]CartItem, error) {
	var cartItems []CartItem
	if err := cartItemRepo.Repo.Where("cart_id=?", cartId).Find(&cartItems).Error; err != nil {
		return cartItems, err
	}
	return cartItems, nil
}
func (cartItemRepo *CartItemRepo) updateTotalAmount(cartId uint64) float64 {
	cartItems, _ := cartItemRepo.FindAllCartItemsByCartId(cartId)
	var totalAmount float64
	for _, cartItem := range cartItems {
		totalAmount += cartItem.TotalPrice
	}
	return totalAmount
}

func (cartItemRepo *CartItemRepo) deleteCart(cartItem CartItem) error {
	if err := cartItemRepo.Repo.Delete(&cartItem).Error; err != nil {
		return err
	}
	return nil
}
func (cartItemRepo *CartItemRepo) RemoveItemFromCart(userId string, productId uint64) error {
	cart, err := cartItemRepo.CartRepo.FindCartByUserId(userId)
	if err != nil {
		return err
	}
	cartItem, err := cartItemRepo.findCartByProductIdAndCartId(cart.Id, productId)
	if err != nil {
		return err
	}
	err = cartItemRepo.deleteCart(cartItem)
	if err != nil {
		return err
	}
	cart.TotalAmount = cartItemRepo.updateTotalAmount(cart.Id)
	err = cartItemRepo.CartRepo.SaveCart(cart)
	if err != nil {
		return err
	}
	return nil
}

func (cartItemRepo *CartItemRepo) DeleteAllCartItem(cartItems []CartItem) error {
	if err := cartItemRepo.Repo.Delete(&cartItems).Error; err != nil {
		return err
	}
	return nil
}
