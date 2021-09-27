package database

import (
	"app/models"
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

// func TestRegisterBossSuccess(t *testing.T) {
// 	config.TestConfig()
// 	config.DB.Migrator().DropTable(&models.Users{})
// 	config.DB.Migrator().AutoMigrate(&models.Users{})
// 	user, err := RegisterBoss(mock_user)
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, "tono", user.Name)
// 		assert.Equal(t, "tono@gmail.com", user.Email)
// 		assert.Equal(t, "123", user.Password)
// 	}
// }

// func TestRegisterBossError(t *testing.T) {
// 	config.TestConfig()
// 	config.DB.Migrator().DropTable(&models.Users{})
// 	_, err := RegisterBoss(mock_user)
// 	assert.Error(t, err)
// }

// func TestLoginBossSuccess(t *testing.T) {
// 	config.TestConfig()
// 	config.DB.Migrator().DropTable(&models.Users{})
// 	config.DB.Migrator().AutoMigrate(&models.Users{})
// 	create_user, _ := RegisterBoss(mock_user)
// 	user, err := LoginBoss(create_user.UserID)
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, "tono", user.Name)
// 		assert.Equal(t, 1, user.UserID)
// 		assert.Equal(t, "tono@gmail.com", user.Email)
// 		assert.Equal(t, "123", user.Password)
// 	}
// }

// func TestLoginBossFail(t *testing.T) {
// 	var user models.Users
// 	var err error
// 	if err = config.DB.Where("user_id= ?", user_id).First(&user).Error; err != nil {
// 		return user, err
// 	}
// 	user.Token, err = middlewares.CreateTokenBoss(int(user.UserID))
// 	if err != nil {
// 		return user, err
// 	}
// 	if err := config.DB.Save(user).Error; err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }
