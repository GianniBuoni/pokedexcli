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
      Name: "map",
      Description: "Displays the next 20 area locations in the Pokemon world",
      Callback: CommandMap,
    },
	}
}

