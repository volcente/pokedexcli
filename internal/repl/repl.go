package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/volcente/pokedexcli/internal/commands"
)

func CleanInput(text string) []string {
	lowercased := strings.ToLower(text)
	return strings.Fields(lowercased)
}

func RunRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	mapConfig := commands.MapConfig{}

	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			return
		}

		input := CleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]

		command, exists := commands.GetCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command.")
			continue
		}

		err := command.Callback(&mapConfig)
		if err != nil {
			fmt.Println(err)
		}
	}
}
