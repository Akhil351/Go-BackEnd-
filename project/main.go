package main

import (
	"net/http"
	"project/web"
	"project/web/handler"
	"project/web/repo"

	"github.com/gorilla/mux"
)

type ProductHandler = handler.ProductHandler
type CartItemHandler = handler.CartItemHandler
type OrderHandler=handler.OrderHandler
type ProductRepo = repo.ProductRepo
type CategoryRepo = repo.CategoryRepo
type CartItemRepo = repo.CartItemRepo
type CartRepo = repo.CartRepo
type OrderItemRepo = repo.OrderItemRepo
type OrderRepo = repo.OrderRepo

func main() {
	web.Init()
	productRepo := ProductRepo{Repo: web.Repo}
	categoryRepo := CategoryRepo{Repo: web.Repo}
	cartRepo := CartRepo{Repo: web.Repo}
	cartItemRepo := CartItemRepo{
		Repo:        web.Repo,
		CartRepo:    &cartRepo,
		ProductRepo: &productRepo,
	}
	orderItemRepo := OrderItemRepo{Repo: web.Repo, CartItemRepo: &cartItemRepo, ProductRepo: &productRepo}
    orderRepo:=OrderRepo{Repo: web.Repo,CartRepo: &cartRepo,OrderItemRepo: &orderItemRepo}
	productHandler := ProductHandler{ProductRepo: &productRepo, CategoryRepo: &categoryRepo}
	CartItemHandler := CartItemHandler{CartItemRepo: &cartItemRepo}
    orderHandler:=OrderHandler{OrderRepo: &orderRepo,ProductRepo: &productRepo}
	r := mux.NewRouter()
	apiPath := "/api/v2"
	r.HandleFunc(apiPath+"/product/addProduct", productHandler.AddProduct).Methods("POST")
	r.HandleFunc(apiPath+"/product/getAllProducts", productHandler.GetAllProducts).Methods("POST")
	r.HandleFunc(apiPath+"/product/updateProduct/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc(apiPath+"/product/deleteProduct/{productId}", productHandler.DeleteProductById).Methods("DELETE")
	r.HandleFunc(apiPath+"/cartItems/add", CartItemHandler.AddItemToCart).Methods("GET")
	r.HandleFunc(apiPath+"/cartItems/remove", CartItemHandler.RemoveItemFromCart).Methods("GET")
	r.HandleFunc(apiPath+"/cartItems/getCart", CartItemHandler.GetCartByUserId).Methods("GET")
	r.HandleFunc(apiPath+"/cartItems/updateCart", CartItemHandler.UpdateQuantity).Methods("PUT")
	r.HandleFunc(apiPath+"/order/placeOrder",orderHandler.PlaceOrder).Methods("GET")
	r.HandleFunc(apiPath+"/order/UserOrder",orderHandler.GetOrder).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", r)
}
