package database

import (
	"app/config"
	"app/middlewares"
	"app/models"
)

func LoginSeller(userId int, password string) (models.Users, error) {
	var seller models.Users
	var err error
	if err = config.DB.Where("user_id = ?", userId, password).First(&seller).Error; err != nil {
		return seller, err
	}
	seller.Token, err = middlewares.CreateTokenSeller(int(seller.UserID))
	if err != nil {
		return seller, err
	}
	if err := config.DB.Save(seller).Error; err != nil {
		return seller, err
	}
	return seller, nil
}

func RegisterSeller(seller models.Users) (models.Users, error) {
	if err := config.DB.Save(&seller).Error; err != nil {
		return seller, err
	}
	return seller, nil
}

func GetSellers() (models.Users, error) {
	var sellers models.Users
	if err := config.DB.Find(&sellers).Error; err != nil {
		return sellers, err
	}
	return sellers, nil
}

func GetOneSeller(id int) (models.Users, error) {
	var seller models.Users
	if err := config.DB.Find(&seller, "user_id=?", id).Error; err != nil {
		return seller, err
	}
	return seller, nil
}
