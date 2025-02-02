package repl

import "fmt"

func helpCommand(mapConfig *Config, commandArgs ...string) error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	fmt.Println("")
	return nil
}
