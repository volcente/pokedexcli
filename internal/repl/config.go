package repl

import "github.com/volcente/pokedexcli/internal/pokeapi"

type Config struct {
	pokeClient       pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func NewConfig(pokeClient pokeapi.Client) Config {
	return Config{
		pokeClient: pokeClient,
	}
}
