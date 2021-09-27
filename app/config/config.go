package config

import (
	"app/models"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var HTTP_PORT int

func InitDb() {
	var err error
	connectionString := "root:toor@tcp(localhost:3306)/pokemon?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitPort() {
	var err error
	HTTP_PORT, err = strconv.Atoi("8080")
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.Transactions{})
	DB.AutoMigrate(&models.Pokemons{})
	DB.AutoMigrate(&models.Users{})
}

func TestConfig() (*gorm.DB, error) {
	var err error
	connectionStringTest := "root:toor@tcp(localhost:3306)/pokemon_testing?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(connectionStringTest), &gorm.Config{})
	if err != nil {
		return DB, err
	}
	return DB, err
}
