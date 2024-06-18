package main

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/azazel-oss/pokedex/internal/pokedex"
)

func commandMapb(config *locationConfig) error {
	if config.Previous == nil {
		return errors.New("you are at the first page")
	}
	locationJson := pokedex.GetPreviousLocations(config.cache, config.Next.String())
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
	} else {
		config.Previous = nil
	}
	return nil
}
