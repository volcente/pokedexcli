package repl

import (
	"errors"
	"fmt"
)

func exploreCommand(config *Config, commandArgs ...string) error {
	if len(commandArgs) == 0 {
		return errors.New("Location name is required. Use map command to list available location areas.")
	}

	locationName := commandArgs[0]
	fmt.Printf("Exploring %s...\n", locationName)
	data, err := config.pokeClient.GetLocationPokemon(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
