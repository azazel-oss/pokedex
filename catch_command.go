package main

import (
	"errors"
	"fmt"

	"github.com/azazel-oss/pokedex/internal/pokedex"
)

func commandCatch(config *locationConfig, words []string) error {
	if len(words) > 2 {
		return errors.New("this commands accepts only one argument")
	}
	if len(words) < 2 {
		return errors.New("this commands needs at least one argument")
	}
	pokemonJson := pokedex.GetPokemonForCatching(config.cache, words[1])
	fmt.Println(pokemonJson.BaseExperience)
	return nil
}
