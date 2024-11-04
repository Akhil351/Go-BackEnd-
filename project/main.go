package main

import (
	"net/http"
	"project/web"
	"project/web/handler"
	"project/web/repo"

	"github.com/gorilla/mux"
)

type ProductHandler = handler.ProductHandler
type ProductRepo = repo.ProductRepo
type CategoryRepo = repo.CategoryRepo

func main() {
	web.Init()
	productRepo := ProductRepo{Repo: web.Repo}
	categoryRepo := CategoryRepo{Repo: web.Repo}
	productHandler := ProductHandler{ProductRepo: productRepo, CategoryRepo: categoryRepo}
	r := mux.NewRouter()
	apiPath := "/api/v2"
	r.HandleFunc(apiPath+"/product/addProduct", productHandler.AddProduct).Methods("POST")
	r.HandleFunc(apiPath+"/product/getAllProducts",productHandler.GetAllProducts).Methods("POST")
	r.HandleFunc(apiPath+"/product/updateProduct/{id}",productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc(apiPath+"/product/deleteProduct/{productId}",productHandler.DeleteProductById).Methods("DELETE")
	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", r)
}
