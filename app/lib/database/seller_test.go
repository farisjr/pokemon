package database

// func TestRegisterSellerSuccess(t *testing.T) {
// 	config.TestConfig()
// 	config.DB.Migrator().DropTable(&models.Users{})
// 	config.DB.Migrator().AutoMigrate(&models.Users{})
// 	user, err := RegisterSeller(mock_user)
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, "tono", user.Name)
// 		assert.Equal(t, "tono@gmail.com", user.Email)
// 		assert.Equal(t, "123", user.Password)
// 	}
// }

// func TestRegisterSellerError(t *testing.T) {
// 	config.TestConfig()
// 	config.DB.Migrator().DropTable(&models.Users{})
// 	_, err := RegisterSeller(mock_user)
// 	assert.Error(t, err)
// }
