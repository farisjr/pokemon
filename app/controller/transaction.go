package controller

import (
	"app/lib/database"
	"app/middlewares"
	"app/models"
	"net/http"

	"github.com/labstack/echo"
)

//Seller Add Transaction
func SellerAddTransaction(c echo.Context) error {
	// auth := AuthorizedSeller(c)
	// if !auth {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
	// }
	addTransaction := models.Transactions{
		Status: "success",
	}
	c.Bind(&addTransaction)
	//Check pokemon from database
	checkPokemon, err := database.GetPokemonFromDatabase(int(addTransaction.PokemonID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get pokemon from database",
		})
	}
	//Checking pokemon stock in database
	if addTransaction.Quantity > checkPokemon.Stock {
		return c.JSON(http.StatusBadRequest, "this pokemon stock is empty")
	}
	//Add transaction to DB
	transaction, err := database.AddTransaction(addTransaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot insert transaction data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add new transaction",
		"data":    transaction,
	})
}

//Get all transaction from database
func BossGetAllTransaction(c echo.Context) error {
	// auth := AuthorizedBoss(c)
	// if !auth {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
	// }
	transactions, err := database.GetAllTransaction()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get transaction list",
		})
	}
	if len(transactions) == 0 {
		return c.JSON(http.StatusInternalServerError, "there is no transaction")
	}
	return c.JSON(http.StatusOK, transactions)
}

func AuthorizedBoss(c echo.Context) bool {
	_, role := middlewares.ExtractToken(c)
	if role != "Boss" {
		return false
	}
	return true
}

func AuthorizedSupplier(c echo.Context) bool {
	_, role := middlewares.ExtractToken(c)
	if role != "Supplier" {
		return false
	}
	return true
}

func AuthorizedSeller(c echo.Context) bool {
	_, role := middlewares.ExtractToken(c)
	if role != "Seller" {
		return false
	}
	return true
}
