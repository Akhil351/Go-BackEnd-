package repo

import (
	"project/web/model"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	Repo *gorm.DB
}
type Category = model.Category

func (categoryRepo *CategoryRepo) AddCategory(categoryName string) (Category, error) {
	existingCategory, err := categoryRepo.FindByCategoryName(categoryName)
	if err == nil {
		return existingCategory, nil
	}
	var category Category
	category.Name = categoryName
	err = categoryRepo.Repo.Save(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (categoryRepo *CategoryRepo) FindByCategoryName(categoryName string) (Category, error) {
	var category Category
	if err := categoryRepo.Repo.Where("name=?", categoryName).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (CategoryRepo *CategoryRepo) FindByCategoryId(categoryId uint64) (string, error) {
	var category Category
	if err := CategoryRepo.Repo.First(&category, categoryId).Error; err != nil {
		return "", err
	}
	return category.Name, nil
}
