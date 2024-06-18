package main

import "net/url"

type cliCommand struct {
	callback         func(*locationConfig) error
	name             string
	description      string
	isConfigRequired bool
}

type locationConfig struct {
	Previous *url.URL
	Next     *url.URL
}

func main() {
	startRepl()
}
