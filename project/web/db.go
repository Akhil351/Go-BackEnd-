package web

import (
	"fmt"
	"os"
	"project/web/model"
	"project/web/repo"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Repo *gorm.DB

type Product = model.Product
type Category = model.Category
type Cart = model.Cart
type CartItem = model.CartItem
type Response = model.Response
type ProductDto = model.ProductDto
type CategoryRepo = repo.CategoryRepo

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error Loading .env file")
		return
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database", err)
		return
	}
	fmt.Println("Connected to the database")
	Repo = database
	Repo.AutoMigrate(&Product{}, &Category{},&Cart{},&CartItem{})

}
