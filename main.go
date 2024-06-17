package main

type cliCommand struct {
	callback    func() error
	name        string
	description string
}

func getPropertiesForCommand(command string) cliCommand {
	commands := map[string]cliCommand{
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
	return commands[command]
}

func main() {
	startRepl()
}
