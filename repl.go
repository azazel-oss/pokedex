package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/azazel-oss/pokedex/internal/pokecache"
)

func startRepl() {
	cache := pokecache.NewCache(5 * time.Second)
	config := initializeConfig(cache)
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		fmt.Println()
		words := sanitizeInput(input.Text())
		commandName := words[0]

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown Command")
			continue
		}
		err := command.callback(&config, words)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func sanitizeInput(input string) []string {
	output := strings.ToLower(input)
	return strings.Fields(output)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "The map command displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore [location-area]",
			description: "Gives you a list of pokemons which can be found in the asked location area",
			callback:    commandExplore,
		},
	}
}

func initializeConfig(cache pokecache.Cache) locationConfig {
	u, err := url.Parse("https://pokeapi.co/api/v2/location-area?offset=0&limit=20")
	if err != nil {
		log.Fatal(err)
	}
	return locationConfig{
		Next:     u,
		Previous: nil,
		cache:    &cache,
	}
}
