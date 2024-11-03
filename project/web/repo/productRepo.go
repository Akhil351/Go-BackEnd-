package repo

import (
	"project/web/model"
	"strconv"

	"gorm.io/gorm"
)

type ProductRepo struct {
	Repo *gorm.DB
}
type Product = model.Product

func (productRepo *ProductRepo) AddProduct(product Product) (Product, error) {
	if err := productRepo.Repo.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}
func (productRepo *ProductRepo) FindProductById(productId uint) (Product, error) {
	var product Product
	if err := productRepo.Repo.First(&product, productId).Error; err != nil {
		return product, err
	}
	return product, nil
}
func (productRepo *ProductRepo) FindAllProducts(searchKey string) ([]Product, error) {
	var products1 []Product
	if searchKey == "" {
		if err := productRepo.Repo.Find(&products1).Error; err != nil {
			return nil, err
		}
		return products1, nil
	}
	key := "%" + searchKey + "%"
	if err := productRepo.Repo.Where("name LIKE ? or brand LIKE ? or description LIKE ? ", key, key, key).Find(&products1).Error; err != nil {
		return nil, err
	}
	var products2 []Product
	num, err := strconv.Atoi(searchKey)
	if err == nil {
		if err := productRepo.Repo.Where("id=? or price=? or inventory=? ", num, num, num).Find(&products2).Error; err != nil {
			return nil, err
		}

	}
	products1 = append(products1, products2...)
	return products1, nil

}
