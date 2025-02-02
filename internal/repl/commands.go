package repl

type cliCommand struct {
	Name        string
	Description string
	Command     func(config *Config, commandArgs ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			Name:        "catch",
			Description: "Attemts to catch a pokemon.",
			Command:     catchCommmand,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays available pokemon in the given location.",
			Command:     exploreCommand,
		},
		"map": {
			Name:        "map",
			Description: "Displays 20 location areas in the Pokemon world.",
			Command:     mapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays 20 previous location areas in the Pokemon world.",
			Command:     mapBackCommand,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message.",
			Command:     helpCommand,
		},
		"exit": {
			Name:        "exit",
			Description: "Exits the Pokedex.",
			Command:     exitCommand,
		},
	}
}
