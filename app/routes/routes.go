package routes

import (
	"app/constants"
	"app/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	e.GET("/seller/pokemons", controller.SellerSearchAskedPokemon)
	e.POST("/seller/transactions", controller.SellerAddTransaction)

	e.GET("/supplier/pokemons/:pokemon_id", controller.SupplierGetPokemonFromPokedex)
	e.POST("/supplier/pokemons/:pokemon_id", controller.SupplierAddPokemonInDatabase)
	e.PUT("/supplier/:pokemon_id", controller.SupplierEditStockPokemon)
	e.DELETE("/supplier/pokemon/:pokemon_id", controller.SupplierDeletePokemonInDatabase)

	e.GET("/boss/transactions", controller.BossGetAllTransaction)

	//------------------ Boss Login & Register Routes ----------------------//
	//e.POST("/boss/login", controller.BossLogin)
	e.POST("/boss/register", controller.BossSignUp)

	//------------------ Supplier Login & Register Routes ----------------------//
	//e.POST("/supplier/login", controller.SupplierLogin)
	e.POST("/supplier/register", controller.SupplierSignUp)

	//------------------ Seller Login & Register Routes ----------------------//
	//e.POST("/seller/login", controller.SellerLogin)
	e.POST("/seller/register", controller.SellerSignUp)

	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// eJwt.GET("/seller/pokemons", controller.SellerSearchAskedPokemon)
	// eJwt.POST("/seller/transactions", controller.SellerAddTransaction)

	// eJwt.GET("/supplier/pokemons", controller.SupplierGetPokemonFromPokedex)
	// eJwt.POST("/supplier/pokemons", controller.SupplierAddPokemonInDatabase)
	// eJwt.PUT("/supplier/:pokemon_id", controller.SupplierEditStockPokemon)
	// eJwt.DELETE("/supplier/pokemon/:pokemon_id", controller.SupplierDeletePokemonInDatabase)

	// eJwt.GET("/boss/transactions", controller.BossGetAllTransaction)

}
