package main

import "fmt"

func commandHelp(_ *locationConfig) error {
	commands := getCommands()
	for _, value := range commands {
		fmt.Println(value.name)
		fmt.Println(value.description)
		fmt.Println("----------------")
	}
	return nil
}
