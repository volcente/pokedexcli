package commands

import (
	"fmt"
	"os"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func HelpCommand() error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	fmt.Println("")
	return nil
}

func ExitCommand() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message.",
			Callback:    HelpCommand,
		},
		"exit": {
			Name:        "exit",
			Description: "Exits the Pokedex.",
			Callback:    ExitCommand,
		},
	}
}
