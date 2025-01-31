package repl

type cliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			Name:        "map",
			Description: "Displays 20 location areas in the Pokemon world.",
			Callback:    mapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays 20 previous location areas in the Pokemon world.",
			Callback:    mapBackCommand,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message.",
			Callback:    helpCommand,
		},
		"exit": {
			Name:        "exit",
			Description: "Exits the Pokedex.",
			Callback:    exitCommand,
		},
	}
}
