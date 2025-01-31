package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			return
		}

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command.")
			continue
		}

		err := command.Callback(config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lowercased := strings.ToLower(text)
	return strings.Fields(lowercased)
}
