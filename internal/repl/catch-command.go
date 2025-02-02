package repl

import (
	"errors"
	"fmt"
	"math/rand"
)

func catchCommmand(config *Config, commandArgs ...string) error {
	if len(commandArgs) == 0 {
		return errors.New("Pokemon name is required! Use explore command to get list of Pokemon for a given location.")
	}

	pokemonName := commandArgs[0]
	if _, alreadyCaught := config.userPokemons[pokemonName]; alreadyCaught {
		fmt.Printf("You already have %s!\n", pokemonName)
		return nil
	}

	data, err := config.pokeClient.GetPokemonDetails(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if !catchPokemon(data.BaseExperience) {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	config.userPokemons[pokemonName] = data

	return nil
}

func catchPokemon(baseExp int) (isCaught bool) {
	catchChance := rand.Intn(100)
	var difficulty int
	switch {
	case baseExp >= 200:
		difficulty = 80
	case baseExp >= 150:
		difficulty = 55
	case baseExp >= 100:
		difficulty = 35
	case baseExp >= 50:
		difficulty = 20
	default:
		difficulty = 10
	}

	return catchChance >= difficulty
}
