package database

import (
	"app/config"
	"app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mock_pokemon = models.Pokemons{
		ID:    1,
		Name:  "ivysaur",
		Stock: 100,
	}
)

func TestCreatePokemonSuccess(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Pokemons{})
	config.DB.Migrator().AutoMigrate(&models.Pokemons{})
	pokemon, err := AddPokemon(mock_pokemon)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, pokemon.ID)
		assert.Equal(t, "ivysaur", pokemon.Name)
		assert.Equal(t, 100, pokemon.Stock)
	}
}

func TestCreatePokemonError(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Pokemons{})
	_, err := AddPokemon(mock_pokemon)
	assert.Error(t, err)
}

func TestGetPokemonByIdSuccess(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Pokemons{})
	config.DB.Migrator().AutoMigrate(&models.Pokemons{})
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	pokemon, err := GetPokemonFromDatabase(int(mock_pokemon.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "ivysaur", pokemon.Name)
		assert.Equal(t, 100, pokemon.Stock)
	}
}

func TestGetPokemonByIdError(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Pokemons{})
	AddPokemon(mock_pokemon)
	_, err := GetPokemonFromDatabase(int(mock_pokemon.ID))
	assert.Error(t, err)
}

func TestGetPokemonByNameSuccess(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Pokemons{})
	config.DB.Migrator().AutoMigrate(&models.Pokemons{})
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	pokemon, err := SearchPokemon(mock_pokemon.Name)
	if assert.NoError(t, err) {
		assert.Equal(t, "ivysaur", pokemon[0].Name)
		assert.Equal(t, 100, pokemon[0].Stock)
	}
}

func TestGetPokemonByNameError(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Pokemons{})
	AddPokemon(mock_pokemon)
	_, err := SearchPokemon(mock_pokemon.Name)
	assert.Error(t, err)
}

func TestDeletePokemonSuccess(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Pokemons{})
	config.DB.Migrator().AutoMigrate(&models.Pokemons{})
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	pokemon, err := DeletePokemon(int(mock_pokemon.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "ivysaur", pokemon.Name)
		assert.Equal(t, 100, pokemon.Stock)
	}
}

func TestDeletePokemonError(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Pokemons{})
	AddPokemon(mock_pokemon)
	_, err := DeletePokemon(int(mock_pokemon.ID))
	assert.Error(t, err)
}

func TestUpdatePokemonSuccess(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Pokemons{})
	config.DB.Migrator().AutoMigrate(&models.Pokemons{})
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	mock_pokemon.Stock = 500
	pokemon, err := EditPokemon(mock_pokemon)
	if assert.NoError(t, err) {
		assert.Equal(t, "ivysaur", pokemon.Name)
		assert.Equal(t, 500, pokemon.Stock)
	}
}

func TestUpdatePokemonError(t *testing.T) {
	config.TestConfig()
	config.DB.Migrator().DropTable(&models.Pokemons{})
	AddPokemon(mock_pokemon)
	mock_pokemon.Stock = 500
	_, err := EditPokemon(mock_pokemon)
	assert.Error(t, err)
}
