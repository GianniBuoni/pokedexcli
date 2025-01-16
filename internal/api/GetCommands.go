package api

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 area locations in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the last 20 area locations in the Pokemon world",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays list of Pokemon in map area. Syntax: explore $AREA_NAME",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempt to catch a Pokemon. Syntax: catch $POKEMON_NAME",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Show info of pokemon registered in Pokedex. Syntax: inspect $POKEMON_NAME",
			Callback:    commandInspect,
		},
	}
}
