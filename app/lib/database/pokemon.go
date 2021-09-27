package database

import (
	"app/config"
	"app/models"
)

func AddPokemon(pokemon models.Pokemons) (models.Pokemons, error) {
	if err := config.DB.Create(&pokemon).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func GetPokemonFromDatabase(pokemon_id int) (models.Pokemons, error) {
	var pokemon models.Pokemons
	if err := config.DB.Find(&pokemon, "id=?", pokemon_id).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func EditPokemon(pokemon models.Pokemons) (models.Pokemons, error) {
	if err := config.DB.Save(&pokemon).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func DeletePokemon(pokemon_id int) (models.Pokemons, error) {
	var pokemon models.Pokemons
	if err := config.DB.First(&pokemon, "id=?", pokemon_id).Error; err != nil {
		return pokemon, err
	}
	if err := config.DB.Delete(&pokemon, "id=?", pokemon_id).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func SearchPokemon(pokemon_name string) ([]models.Pokemons, error) {
	var pokemon []models.Pokemons
	name := ("%" + pokemon_name + "%")
	if err := config.DB.Find(&pokemon, "name LIKE ?", name).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}
