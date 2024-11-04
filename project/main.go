package main

import (
	"net/http"
	"project/web"
	"project/web/handler"
	"project/web/repo"

	"github.com/gorilla/mux"
)

type ProductHandler = handler.ProductHandler
type CartItemHandler=handler.CartItemHandler
type ProductRepo = repo.ProductRepo
type CategoryRepo = repo.CategoryRepo
type CartItemRepo=repo.CartItemRepo
type CartRepo=repo.CartRepo
func main() {
	web.Init()
	productRepo := ProductRepo{Repo: web.Repo}
	categoryRepo := CategoryRepo{Repo: web.Repo}
	cartRepo:=CartRepo{Repo: web.Repo}
	cartItemRepo:=CartItemRepo{
		Repo: web.Repo,
		CartRepo: &cartRepo,
		ProductRepo: &productRepo,
	}
	productHandler := ProductHandler{ProductRepo: &productRepo, CategoryRepo: &categoryRepo}
	CartItemHandler:=CartItemHandler{CartItemRepo: &cartItemRepo}

	r := mux.NewRouter()
	apiPath := "/api/v2"
	r.HandleFunc(apiPath+"/product/addProduct", productHandler.AddProduct).Methods("POST")
	r.HandleFunc(apiPath+"/product/getAllProducts", productHandler.GetAllProducts).Methods("POST")
	r.HandleFunc(apiPath+"/product/updateProduct/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc(apiPath+"/product/deleteProduct/{productId}", productHandler.DeleteProductById).Methods("DELETE")
	r.HandleFunc(apiPath+"/cartItems/add",CartItemHandler.AddItemToCart).Methods("GET")
	r.HandleFunc(apiPath+"/cartItems/remove",CartItemHandler.RemoveItemFromCart).Methods("GET")
	r.HandleFunc(apiPath+"/cartItems/getCart",CartItemHandler.GetCartByUserId).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", r)
}
