package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
)

func CreateResponse(w http.ResponseWriter, err error, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	status := "Success"
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
		status = "Failed"
		w.WriteHeader(http.StatusBadRequest)
	}
	response := Response{
		Status:    status,
		TimeStamp: time.Now(),
		Data:      data,
		Error:     errMsg,
	}
	json.NewEncoder(w).Encode(response)

}

func ConvertToDto[S any,T any](source S,destination T)(T){
	copier.Copy(&destination,&source)
	return destination
}

func CategoryNameToProductDto(productDto *ProductDto,categoryRepo *CategoryRepo,product Product){
	categoryName, _ := categoryRepo.FindByCategoryId(product.CategoryId)
	productDto.CategoryName = categoryName
}

func AssignProductName(cartItemDto *CartItemDto,productRepo *ProductRepo,productId uint64){
	product,_:=productRepo.FindProductById(productId)
	cartItemDto.ProductName=product.Name
}

func AssignProductNameToOrder(orderItemDto *OrderItemDto,productRepo *ProductRepo,productId uint64){
	product,_:=productRepo.FindProductById(productId)
	orderItemDto.ProductName=product.Name
}

func ProductRequestToProduct(product Product, productRequest ProductDto, categoryRepo *CategoryRepo) Product {

	if productRequest.Name != "" {
		product.Name = productRequest.Name
	}
	if productRequest.Brand != "" {
		product.Brand = productRequest.Brand
	}
	if productRequest.Description != "" {
		product.Description = productRequest.Description
	}
	if productRequest.Inventory > 0 {
		product.Inventory = uint(productRequest.Inventory)
	}
	if productRequest.Price > 0 {
		product.Price = productRequest.Price
	}
	if productRequest.CategoryName != "" {
		category, _ := categoryRepo.AddCategory(productRequest.CategoryName)
		product.CategoryId = category.Id
	}
	return product
}

func ValidateProductRequest(productRequest ProductDto) error {
	if productRequest.Name == "" {
		return fmt.Errorf(" name is not provided")
	}
	if productRequest.Brand == "" {
		return fmt.Errorf("brand  is not provided")
	}
	if productRequest.Description == "" {
		return fmt.Errorf("description  is not provided")
	}
	if productRequest.Inventory <= 0 {
		return fmt.Errorf("inventory  should be greater than zero")
	}
	if productRequest.Price <= 0 {
		return fmt.Errorf("price  should be greater than zero")
	}
	if productRequest.CategoryName == "" {
		return fmt.Errorf("categoryName  is not provided")
	}
	return nil

}
