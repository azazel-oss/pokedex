package main

import (
	"errors"
	"fmt"

	"github.com/azazel-oss/pokedex/internal/pokedex"
)

func commandExplore(config *locationConfig, words []string) error {
	if len(words) > 2 {
		return errors.New("this commands accepts only one argument")
	}
	if len(words) < 2 {
		return errors.New("this commands needs at least one argument")
	}
	locationAreaJson := pokedex.GetPokemonsByLocationArea(config.cache, words[1])
	for _, pokemon := range locationAreaJson.PokemonEncounters {
		fmt.Print("- ")
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
