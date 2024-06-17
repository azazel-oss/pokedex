package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
		err := command.callback()
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
	}
}
