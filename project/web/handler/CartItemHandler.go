package handler

import (
	"net/http"
	"project/web"
	"project/web/repo"
	"strconv"
)

type CartItemRepo = repo.CartItemRepo
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
	err = handler.CartItemRepo.AddItemToCart(userId, uint64(productId), quantity)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	web.CreateResponse(w,nil,"CartItem Added Succssfully")
}
