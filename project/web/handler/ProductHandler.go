package handler

import (
	"encoding/json"
	"net/http"
	"project/web"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	ProductRepo  *ProductRepo
	CategoryRepo *CategoryRepo
}

func (handler *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var productRequest ProductDto
	if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	err := web.ValidateProductRequest(productRequest)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	var product Product
	product = web.ProductRequestToProduct(product, productRequest, handler.CategoryRepo)
	newProduct, err := handler.ProductRepo.AddProduct(product)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	productRequest.Id = newProduct.Id
	web.CreateResponse(w, nil, productRequest)

}
func (handler *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	products, err := handler.ProductRepo.FindAllProducts(request.SearchKey)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	var productsDto []ProductDto
	for _, product := range products {
		productDto := web.ConvertToDto(product, ProductDto{})
		web.CategoryNameToProductDto(&productDto, handler.CategoryRepo, product)
		productsDto = append(productsDto, productDto)
	}
	web.CreateResponse(w, nil, productsDto)
}

func (handler *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var productRequest ProductDto
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	var product Product
	product, err = handler.ProductRepo.FindProductById(uint64(id))
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	product = web.ProductRequestToProduct(product, productRequest, handler.CategoryRepo)
	updatedProduct, err := handler.ProductRepo.AddProduct(product)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	productDto := web.ConvertToDto(updatedProduct, ProductDto{})
	web.CategoryNameToProductDto(&productDto, handler.CategoryRepo, updatedProduct)
	web.CreateResponse(w, nil, productDto)

}

func (handler *ProductHandler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["productId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	err = handler.ProductRepo.DeleteProduct(uint64(id))
	if err != nil {
		web.CreateResponse(w, err, nil)
		return
	}
	web.CreateResponse(w, nil, "Product Delete Successfully")

}
