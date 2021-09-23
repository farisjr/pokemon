package database

import (
	"app/config"
	"app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionSuccess(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Users{})
	config.DB.Migrator().AutoMigrate(&models.Users{})
	config.DB.Migrator().DropTable(&models.Transactions{})
	config.DB.Migrator().AutoMigrate(&models.Transactions{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transactions{
		TotalPrice: 10000,
		PokemonID:  uint(1),
		SellerID:   uint(1),
		Status:     "Success",
	}
	transaction, err := AddTransaction(mock_transaction)
	if assert.NoError(t, err) {
		assert.Equal(t, 10000, transaction.TotalPrice)
		assert.Equal(t, "Success", transaction.Status)
		assert.Equal(t, uint(1), transaction.PokemonID)
		assert.Equal(t, uint(1), transaction.SellerID)
	}
}

func TestCreateTransactionError(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Users{})
	config.DB.Migrator().AutoMigrate(&models.Users{})
	config.DB.Migrator().DropTable(&models.Transactions{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transactions{
		TotalPrice: 10000,
		Status:     "Success",
	}
	_, err := AddTransaction(mock_transaction)
	assert.Error(t, err)
}

func TestGetAllTransactionSuccess(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Users{})
	config.DB.Migrator().AutoMigrate(&models.Users{})
	config.DB.Migrator().DropTable(&models.Transactions{})
	config.DB.Migrator().AutoMigrate(&models.Transactions{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transactions{
		Status: "Success",
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	transaction, err := GetAllTransaction()
	if assert.NoError(t, err) {
		assert.Equal(t, 10000, transaction[0].TotalPrice)
		assert.Equal(t, "Success", transaction[0].Status)
	}
}

func TestGetAllTransactionFailed(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Users{})
	config.DB.Migrator().AutoMigrate(&models.Users{})
	config.DB.Migrator().DropTable(&models.Transactions{})
	config.DB.Migrator().AutoMigrate(&models.Transactions{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transactions{
		PokemonID: 1,
		Status:    "Cancelled",
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	transaction, err := GetAllTransaction()
	if assert.NoError(t, err) {
		assert.Equal(t, 10000, transaction[0].TotalPrice)
		assert.Equal(t, "Cancelled", transaction[0].Status)
	}
}
