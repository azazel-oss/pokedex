package main

import (
	"fmt"
	"net/url"

	"github.com/azazel-oss/pokedex/internal/pokedex"
)

func commandMap(config *locationConfig, _ []string) error {
	locationJson := pokedex.GetNextLocations(config.cache, config.Next.String())
	for _, location := range locationJson.Results {
		fmt.Println(location.Name)
	}
	if locationJson.Next != nil {
		u, err := url.Parse(*locationJson.Next)
		if err != nil {
			return err
		}
		config.Next = u
	} else {
		config.Next = nil
	}

	if locationJson.Previous != nil {
		p, err := url.Parse(*locationJson.Previous)
		if err != nil {
			return err
		}
		config.Previous = p
	}
	return nil
}
