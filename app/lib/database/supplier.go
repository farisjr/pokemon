package database

import (
	"app/config"
	"app/middlewares"
	"app/models"
)

func LoginSupplier(userId int, password string) (models.Users, error) {
	var supplier models.Users
	var err error
	if err = config.DB.Where("user_id = ?", userId, password).First(&supplier).Error; err != nil {
		return supplier, err
	}
	supplier.Token, err = middlewares.CreateTokenBoss(int(supplier.UserID))
	if err != nil {
		return supplier, err
	}
	if err := config.DB.Save(supplier).Error; err != nil {
		return supplier, err
	}
	return supplier, nil
}

func RegisterSupplier(supplier models.Users) (models.Users, error) {
	if err := config.DB.Save(&supplier).Error; err != nil {
		return supplier, err
	}
	return supplier, nil
}

func GetSupplier() (models.Users, error) {
	var supplier models.Users
	if err := config.DB.Find(&supplier).Error; err != nil {
		return supplier, err
	}
	return supplier, nil
}

func GetOneSupplier(id int) (models.Users, error) {
	var supplier models.Users
	if err := config.DB.Find(&supplier, "user_id=?", id).Error; err != nil {
		return supplier, err
	}
	return supplier, nil
}
