package database

import (
	"app/config"
	"app/middlewares"
	"app/models"
)

func LoginBoss(userId int, password string) (models.Users, error) {
	var boss models.Users
	var err error
	if err = config.DB.Where("user_id = ?", userId, password).First(&boss).Error; err != nil {
		return boss, err
	}
	boss.Token, err = middlewares.CreateTokenBoss(int(boss.UserID))
	if err != nil {
		return boss, err
	}
	if err := config.DB.Save(boss).Error; err != nil {
		return boss, err
	}
	return boss, nil
}

func RegisterBoss(boss models.Users) (models.Users, error) {
	if err := config.DB.Save(&boss).Error; err != nil {
		return boss, err
	}
	return boss, nil
}

func GetBoss() (models.Users, error) {
	var boss models.Users
	if err := config.DB.Find(&boss).Error; err != nil {
		return boss, err
	}
	return boss, nil
}

func GetOneBoss(id int) (models.Users, error) {
	var boss models.Users
	if err := config.DB.Find(&boss, "user_id=?", id).Error; err != nil {
		return boss, err
	}
	return boss, nil
}

func UpdateBoss(boss models.Users) (models.Users, error) {
	if err := config.DB.Save(&boss).Error; err != nil {
		return boss, err
	}
	return boss, nil
}
