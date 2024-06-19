package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/azazel-oss/pokedex/internal/pokedex"
)

func commandCatch(config *locationConfig, words []string) error {
	if len(words) > 2 {
		return errors.New("this commands accepts only one argument")
	}
	if len(words) < 2 {
		return errors.New("this commands needs at least one argument")
	}
	pokemonJson, err := pokedex.GetPokemonForCatching(config.cache, words[1])
	if err != nil {
		return err
	}
	random := rand.Intn(pokemonJson.BaseExperience)
	if random > pokemonJson.BaseExperience-30 {
		fmt.Println(pokemonJson.Name + " has been caught. Congratulations.")
		config.userPokedex[pokemonJson.Name] = Pokemon{
			Name:           pokemonJson.Name,
			BaseExperience: pokemonJson.BaseExperience,
			Id:             pokemonJson.ID,
		}
	} else {
		fmt.Println(pokemonJson.Name + " has escaped and fled the scene")
	}
	return nil
}
