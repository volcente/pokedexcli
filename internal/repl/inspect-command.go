package repl

import (
	"errors"
	"fmt"
)

func inspectCommand(config *Config, commandArgs ...string) error {
	if len(commandArgs) == 0 {
		return errors.New("Pokemon name is required.")
	}

	pokemonName := commandArgs[0]
	pokemon, isCaught := config.userPokemons[pokemonName]
	if !isCaught {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println()
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, variant := range pokemon.Types {
		fmt.Printf("  - %s\n", variant.Type.Name)
	}
	fmt.Println()

	return nil
}
