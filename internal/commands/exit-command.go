package commands

import (
	"fmt"
	"os"
)

func exitCommand(mapConfig *MapConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
