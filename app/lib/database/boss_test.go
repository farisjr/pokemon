package database

import (
	"app/config"
	"app/models"
	"testing"
)

var (
	mock_user_boss = models.Users{
		Name:     "Tono",
		Email:    "tono@gmail.com",
		Password: "123",
		Role:     "Boss",
	}
	mock_boss_login = models.Users{
		Email:    "tono@gmail.com",
		Password: "123",
	}
)

func TestLoginSuccess(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Users{})
	config.DB.Migrator().AutoMigrate(&models.Users{})
	//create_user, _ := RegisterBoss(mock_user)
	//user, err := LoginBoss(create_user.UserID)
	// if assert.NoError(t, err) {
	// 	assert.Equal(t, "1", user.UserID)
	// 	assert.Equal(t, "1", user.Password)
	// }
}
