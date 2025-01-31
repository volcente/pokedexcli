package repl

import (
	"fmt"
	"os"
)

func exitCommand(mapConfig *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
