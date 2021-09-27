package controller

import (
	"app/lib/database"
	"app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Consume pokedex API to check pokemon
func SupplierGetPokemonFromPokedex(c echo.Context) error {
	//auth := AuthorizedSupplier(c)
	// if !auth {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
	// }
	pokemon_id, err := strconv.Atoi(c.Param("pokemon_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid pokemon id",
		})
	}
	pokedex_url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", pokemon_id) //Consuming Pokedex API
	response, err := http.Get(pokedex_url)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}
	response_data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var pokemon models.Pokedexs
	json.Unmarshal(response_data, &pokemon)
	return c.JSON(http.StatusOK, pokemon)
}

// Add Supplier add Pokemon to Database
func SupplierAddPokemonInDatabase(c echo.Context) error {
	// auth := AuthorizedSupplier(c)
	// if !auth {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
	// }
	pokemon_id, err := strconv.Atoi(c.Param("pokemon_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid pokemon id",
		})
	}
	//Consuming Pokedex API
	pokedex_url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", pokemon_id)
	response, err := http.Get(pokedex_url)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}
	response_data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var pokedex models.Pokedexs
	json.Unmarshal(response_data, &pokedex)
	pokemon := models.Pokemons{
		ID:   pokedex.ID,
		Name: pokedex.Name,
	}
	c.Bind(&pokemon)
	addPokemonDb, err := database.AddPokemon(pokemon) //Add Pokemon
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, addPokemonDb)
}

//Supplier Edit Existing Pokemon in Database
func SupplierEditStockPokemon(c echo.Context) error {
	// auth := AuthorizedSupplier(c)
	// if !auth {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
	// }
	pokemon_id, err := strconv.Atoi(c.Param("pokemon_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid pokemon id",
		})
	}
	pokemon, err := database.GetPokemonFromDatabase(pokemon_id) //Get Pokemon from Database
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get pokemon from database",
		})
	}
	c.Bind(&pokemon)
	editted_pokemon, err := database.EditPokemon(pokemon) //Edit Pokemon
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit pokemon's stock",
		})
	}
	return c.JSON(http.StatusOK, editted_pokemon)
}

//Supplier Delete Pokemon from Database
func SupplierDeletePokemonInDatabase(c echo.Context) error {
	// auth := AuthorizedSupplier(c)
	// if !auth {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
	// }
	pokemon_id, err := strconv.Atoi(c.Param("pokemon_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid pokemon id",
		})
	}
	deletted_pokemon, err := database.DeletePokemon(pokemon_id) //Delete Pokemon
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, deletted_pokemon)
}

// Seller get pokemon by customer request
func SellerSearchAskedPokemon(c echo.Context) error {
	// auth := AuthorizedSeller(c)
	// if !auth {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
	// }
	pokemon_name := c.QueryParam("pokemon_name")
	pokemons, err := database.SearchPokemon(pokemon_name) //Search Pokemon
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, pokemons)
}
