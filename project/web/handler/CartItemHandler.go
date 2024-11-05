package handler

import (
	"net/http"
	"project/web"
	"strconv"
)

type CartItemHandler struct {
	CartItemRepo *CartItemRepo
}

func (handler *CartItemHandler) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	productIdstr := r.URL.Query().Get("productId")
	productId, err := strconv.Atoi(productIdstr)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	quantityStr := r.URL.Query().Get("quantity")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	err = handler.CartItemRepo.AddItemToCart(userId, uint64(productId), quantity,"Add")
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	web.CreateResponse(w, nil, "CartItem Added Successfully")
}
func (handler *CartItemHandler) RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	productIdstr := r.URL.Query().Get("productId")
	productId, err := strconv.Atoi(productIdstr)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	err = handler.CartItemRepo.RemoveItemFromCart(userId, uint64(productId))
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	web.CreateResponse(w, nil, "CartItem Removed Successfully")
}

func (handler *CartItemHandler) GetCartByUserId(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	cart, err := handler.CartItemRepo.CartRepo.FindCartByUserId(userId)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	cartItems, err := handler.CartItemRepo.FindAllCartItemsByCartId(cart.Id)
	var cartItemsDto []CartItemDto
	for _, cartItem := range cartItems {
		cartItemDto := web.ConvertToDto(cartItem, CartItemDto{})
		web.ProductNameToCartItemDto(&cartItemDto, handler.CartItemRepo.ProductRepo, cartItem.ProductId)
		cartItemsDto = append(cartItemsDto, cartItemDto)
	}
	cartDto := web.ConvertToDto(cart, CartDto{})
	cartDto.CartItems = cartItemsDto
	web.CreateResponse(w, nil, cartDto)

}
func(handler *CartItemHandler) UpdateQuantity(w http.ResponseWriter,r *http.Request){
	userId := r.URL.Query().Get("userId")
	productIdstr := r.URL.Query().Get("productId")
	productId, err := strconv.Atoi(productIdstr)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	quantityStr := r.URL.Query().Get("quantity")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	err=handler.CartItemRepo.AddItemToCart(userId,uint64(productId),quantity,"Update")
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	web.CreateResponse(w, nil, "Quantity Updated Successfully")


}
