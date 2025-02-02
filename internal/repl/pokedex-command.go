package repl

import "fmt"

func pokedexCommand(config *Config, commandArgs ...string) error {
	fmt.Println("Your Pokedex:")

	userPokemons := config.userPokemons
	if len(userPokemons) == 0 {
		fmt.Println("You've not caught a Pokemon yet. Try using catch command with the name of pokemon you want to catch!")
		return nil
	}

	for _, pokemon := range userPokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
