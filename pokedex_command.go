package main

import "fmt"

func commandPokedex(config *locationConfig, _ []string) error {
	pokedex := config.userPokedex
	fmt.Println("Your pokedex: ")

	for key := range pokedex {
		fmt.Println(" - ", key)
	}

	return nil
}
