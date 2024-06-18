package main

import (
	"net/url"

	"github.com/azazel-oss/pokedex/internal/pokecache"
)

type cliCommand struct {
	callback         func(*locationConfig) error
	name             string
	description      string
	isConfigRequired bool
}

type locationConfig struct {
	Previous *url.URL
	Next     *url.URL
	cache    *pokecache.Cache
}

func main() {
	startRepl()
}
