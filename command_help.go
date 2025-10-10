package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("Hello adventurer!")
	fmt.Println("This is the help menu, for your pokedex.")
	fmt.Println("Available commands are:")
	availableCommands := getCommands()
	fmt.Println()
	for _, c := range availableCommands {
		fmt.Printf("- %s: %s \n", c.name, c.description)

	}
	return nil
}
