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
	auth := AuthorizedSeller(c)
	if !auth {
		return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
	}
	addTransaction := models.Transactions{}
	c.Bind(&addTransaction)
	transaction, err := database.AddTransaction(addTransaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot insert data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add new transaction",
		"data":    transaction,
	})
}

//Get List of Transaction from Database
func BossGetAllTransaction(c echo.Context) error {
	auth := AuthorizedBoss(c)
	if !auth {
		return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
	}
	transactions, err := database.GetListofTransaction()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get transaction list",
		})
	}
	//Handling error
	if len(transactions) == 0 {
		return c.JSON(http.StatusInternalServerError, "there's no transaction")
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
