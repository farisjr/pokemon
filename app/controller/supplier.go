package controller

import (
	"app/lib/database"
	"app/models"
	"net/http"

	"github.com/labstack/echo"
)

//Register supplier controller for supplier registration
func SupplierSignUp(c echo.Context) error {
	input := models.Users{}
	c.Bind(&input)
	if input.UserID == 0 || input.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please fill userid and password correctly",
		})
	}
	if same, _ := database.CheckSameUserId(input.UserID); same == true {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "userid already used",
		})
	}
	addSupplier := models.Users{}
	addSupplier.UserID = input.UserID
	//addSupplier.Password = database.OurEncrypt(input.Password)
	addSupplier.Role = "Supplier"
	c.Bind(&addSupplier)
	supplier, err := database.RegisterSeller(addSupplier)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	mapSupplier := map[string]interface{}{
		"User ID": supplier.UserID,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new supplier",
		"data":    mapSupplier,
	})
}

//Login for supplier with matching userid and password
// func SupplierLogin(c echo.Context) error {
// 	input := models.Users{}
// 	c.Bind(&input)
// 	loginSupplier, err := database.LoginSupplier(input.UserID, input.Password)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	mapLoginSupplier := map[string]interface{}{
// 		"UserID": loginSupplier.UserID,
// 		"Token":  loginSupplier.Token,
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "succes login",
// 		"data":    mapLoginSupplier,
// 	})
// }

//Authorization Supplier
// func SupplierAutorize(supplierId int, c echo.Context) error {
// 	authSupplier, err := database.GetOneSeller(supplierId)
// 	LoggedInSupplier, role := middlewares.ExtractToken(c)
// 	if LoggedInSupplier != supplierId || string(authSupplier.Role) != role || err != nil || authSupplier.Role != "Supplier" {
// 		return echo.NewHTTPError(http.StatusUnauthorized, "This user does not have access")
// 	}
// 	return nil
// }
