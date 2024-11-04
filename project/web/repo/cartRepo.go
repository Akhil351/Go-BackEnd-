package repo

import (
	"project/web/model"

	"gorm.io/gorm"
)

type Cart=model.Cart
type CartRepo struct{
  Repo *gorm.DB
}

func (cartRepo *CartRepo) FindCartByUserId(userId string) (Cart,error){
	var cart Cart
	if err:=cartRepo.Repo.Where("user_id=?",userId).First(&cart).Error; err!=nil{
		return cart,err
	}
	return cart,nil
}

func(cartRepo *CartRepo) SaveCart(cart Cart) (error){
	if err:=cartRepo.Repo.Save(&cart).Error ; err!=nil{
		return err
	}
	return nil
}