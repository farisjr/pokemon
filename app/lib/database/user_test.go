package database

import (
	"app/config"
	"app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mock_user = models.Users{
		Name:     "tono",
		Email:    "tono@gmail.com",
		Password: "123",
		Role:     "Boss",
	}
)

func TestCheckEmailError(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Users{})
	RegisterSeller(mock_user)
	_, err := CheckEmail("Tono@gmail.com")
	assert.Error(t, err)
}

func TestCheckEmailSameSuccess(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Users{})
	config.DB.Migrator().AutoMigrate(&models.Users{})
	RegisterSeller(mock_user)
	same, err := CheckEmail("tono@gmail.com")
	if assert.NoError(t, err) {
		assert.Equal(t, true, same)
	}
}

func TestCheckUserIdError(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Users{})
	RegisterSeller(mock_user)
	_, err := CheckSameUserId(1)
	assert.Error(t, err)
}

func TestCheckUserIdSuccess(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Users{})
	config.DB.Migrator().AutoMigrate(&models.Users{})
	RegisterSeller(mock_user)
	same, err := CheckSameUserId(1)
	if assert.NoError(t, err) {
		assert.Equal(t, true, same)
	}
}
