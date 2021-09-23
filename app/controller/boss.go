package controller

import (
	"app/lib/database"
	"app/middlewares"
	"app/models"
	"net/http"

	"github.com/labstack/echo"
)

//Register boss controller for boss registration
func BossSignUp(c echo.Context) error {
	input := models.Users{}
	c.Bind(&input)
	if input.UserID == 0 || input.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please fill user id and password correctly",
		})
	}
	if same, _ := database.CheckSameUserId(input.UserID); same == true {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "user id already used",
		})
	}
	addBoss := models.Users{}
	addBoss.UserID = input.UserID
	addBoss.Password = database.OurEncrypt(input.Password)
	addBoss.Role = "Boss"
	c.Bind(&addBoss)
	seller, err := database.RegisterSeller(addBoss)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	mapBoss := map[string]interface{}{
		"User ID": seller.UserID,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new boss",
		"data":    mapBoss,
	})
}

//Login for seller with matching userid and password
func BossLogin(c echo.Context) error {
	input := models.Users{}
	c.Bind(&input)
	loginBoss, err := database.LoginBoss(input.UserID, input.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	mapLoginBoss := map[string]interface{}{
		"UserID": loginBoss.UserID,
		"Token":  loginBoss.Token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"data":    mapLoginBoss,
	})
}

//Authorization Boss
func BossAutorize(bossId int, c echo.Context) error {
	authBoss, err := database.GetOneBoss(bossId)
	LoggedInBoss, role := middlewares.ExtractToken(c)
	if LoggedInBoss != bossId || string(authBoss.Role) != role || err != nil || authBoss.Role != "Boss" {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user does not have access")
	}
	return nil
}
