package database

import (
	"app/config"
	"app/models"
)

func AddTransaction(transaction models.Transactions) (models.Transactions, error) {
	if err := config.DB.Save(&transaction).Error; err != nil {
		return transaction, err
	}
	// pokemon.Stock -= transaction.Quantity //Auto Update Stock
	if err := config.DB.Save(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func GetListofTransaction() ([]models.Transactions, error) {
	var transactions []models.Transactions
	if err := config.DB.Preload("Pokemon").Preload("User").Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}

func GetTransactionById(transaction_id int) (models.Transactions, error) {
	var transaction models.Transactions
	if err := config.DB.Preload("Pokemon").Preload("User").First(&transaction, "id=?", transaction_id).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
