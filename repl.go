package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func startRepl() {
	config := initializeConfig()
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
		err := command.callback(&config)
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
			name:             "map",
			description:      "The map command displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on.",
			callback:         commandMap,
			isConfigRequired: true,
		},
		"mapb": {
			name:             "mapb",
			description:      "Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
			callback:         commandMapb,
			isConfigRequired: true,
		},
	}
}

func initializeConfig() locationConfig {
	u, err := url.Parse("https://pokeapi.co/api/v2/location")
	if err != nil {
		log.Fatal(err)
	}
	return locationConfig{
		Next:     u,
		Previous: nil,
	}
}
