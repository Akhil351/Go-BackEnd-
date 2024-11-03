package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CreateResponse(w http.ResponseWriter, err error, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	status := "Success"
	errMsg:=""
	if err != nil {
		errMsg=err.Error()
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

func ProductToProductDto(product Product, categoryRepo CategoryRepo) ProductDto {
	var productDto ProductDto
	productDto.Name = product.Name
	productDto.Brand = product.Brand
	productDto.Description = product.Description
	productDto.Inventory = int(product.Inventory)
	productDto.Price = product.Price
	productDto.Id = int(product.Id)
	categoryName, _ := categoryRepo.FindByCategoryId(product.CategoryId)
	productDto.CategoryName = categoryName
	return productDto
}

func ProductRequestToProduct( product Product,productRequest ProductDto, categoryRepo CategoryRepo) (Product) {
	
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
	if productRequest.Inventory == 0 {
		return fmt.Errorf("inventory  is not provided")
	}
	if productRequest.Price == 0 {
		return fmt.Errorf("price  is not provided")
	}
	if productRequest.CategoryName == "" {
		return fmt.Errorf("categoryName  is not provided")
	}
	return nil

}
