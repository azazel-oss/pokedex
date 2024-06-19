package main

import (
	"net/url"

	"github.com/azazel-oss/pokedex/internal/pokecache"
)

type cliCommand struct {
	callback    func(*locationConfig, []string) error
	name        string
	description string
}

type locationConfig struct {
	Previous    *url.URL
	Next        *url.URL
	cache       *pokecache.Cache
	userPokedex map[string]Pokemon
}

func main() {
	startRepl()
}
