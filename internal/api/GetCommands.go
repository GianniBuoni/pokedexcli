package api

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 area locations in the Pokemon world",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the last 20 area locations in the Pokemon world",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays list of Pokemon in map area. Syntax: explore $AREA_NAME",
			Callback:    explore,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempt to catch a Pokemon. Syntax: catch $POKEMON_NAME",
			Callback:    catch,
		},
	}
}
