package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *locationConfig, words []string) error {
	if len(words) > 2 {
		return errors.New("this commands accepts only one argument")
	}
	if len(words) < 2 {
		return errors.New("this commands needs at least one argument")
	}
	pokemonName := words[1]
	pokemon, err := getPokemonDetails(config.userPokedex, pokemonName)
	if err != nil {
		return err
	}
	printPokemonDetails(pokemon)
	return nil
}

func getPokemonDetails(pokedex map[string]Pokemon, pokemon string) (Pokemon, error) {
	if pokemon, ok := pokedex[pokemon]; ok {
		return pokemon, nil
	}
	return Pokemon{}, errors.New("you haven't caught this pokemon yet")
}

func printPokemonDetails(pokemon Pokemon) {
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Base Experience: ", pokemon.BaseExperience)
	fmt.Println("Id: ", pokemon.Id)
}
