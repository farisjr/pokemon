package controller

import (
	"app/lib/database"
	"app/models"
	"net/http"

	"github.com/labstack/echo"
)

//Register seller controller for seller registration
func SellerSignUp(c echo.Context) error {
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
	addSeller := models.Users{}
	addSeller.UserID = input.UserID
	//addSeller.Password = database.OurEncrypt(input.Password)
	addSeller.Role = "Seller"
	c.Bind(&addSeller)
	seller, err := database.RegisterSeller(addSeller)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	mapSeller := map[string]interface{}{
		"User ID": seller.UserID,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new seller",
		"data":    mapSeller,
	})
}

// //Login for seller with matching userid and password
// func SellerLogin(c echo.Context) error {
// 	input := models.Users{}
// 	c.Bind(&input)
// 	loginSeller, err := database.LoginSeller(input.UserID, input.Password)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	mapLoginSeller := map[string]interface{}{
// 		"UserID": loginSeller.UserID,
// 		"Token":  loginSeller.Token,
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "succes login",
// 		"data":    mapLoginSeller,
// 	})
// }

//Authorization Seller
// func SellerAutorize(sellerId int, c echo.Context) error {
// 	authSeller, err := database.GetOneSeller(sellerId)
// 	LoggedInSeller, role := middlewares.ExtractToken(c)
// 	if LoggedInSeller != sellerId || string(authSeller.Role) != role || err != nil || authSeller.Role != "Seller" {
// 		return echo.NewHTTPError(http.StatusUnauthorized, "This user does not have access")
// 	}
// 	return nil
// }
