package repl

import "github.com/volcente/pokedexcli/internal/pokeapi"

type Config struct {
	userPokemons     map[string]pokeapi.Pokemon
	pokeClient       pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func NewConfig(pokeClient pokeapi.Client) Config {
	return Config{
		pokeClient:   pokeClient,
		userPokemons: map[string]pokeapi.Pokemon{},
	}
}
