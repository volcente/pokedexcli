package commands

import "fmt"

func helpCommand(mapConfig *MapConfig) error {
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
