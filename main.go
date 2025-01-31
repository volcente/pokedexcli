package main

import (
	"time"

	"github.com/volcente/pokedexcli/internal/pokeapi"
	"github.com/volcente/pokedexcli/internal/repl"
)

func main() {
	apiClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	config := repl.NewConfig(apiClient)

	repl.RunRepl(&config)
}
