package database

import (
	"app/config"
	"app/models"
)

func CheckEmail(email string) (bool, error) {
	var user models.Users
	if err := config.DB.Model(&user).Where("email=?", email).First(&user).Error; err != nil {
		return false, err
	}
	if user.Email == email {
		return true, nil
	} else {
		return false, nil
	}
}

func CheckSameUserId(userid uint) (bool, error) {
	var user models.Users
	if err := config.DB.Raw("select * from users where user_id = ?", userid).Scan(&user).Error; err != nil {
		return true, err
	}
	if user.UserID == userid {
		return true, nil
	}
	return false, nil
}

// func OurEncrypt(plain string) string {
// 	bytePlain := []byte(plain)
// 	hashed, _ := bcrypt.GenerateFromPassword(bytePlain, bcrypt.MinCost)
// 	return string(hashed)
// }
