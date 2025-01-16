# Welcome to the Pokedex!

This is my first attempt to make a terminal app in Go.
I'm still learning the basics of making HTTP requests, setting up a cache, and decoding JSON.

## Usage:

- `help`: Displays a help message
- `map`: Displays the next 20 area locations in the Pokemon world
- `mapb`: Displays the last 20 area locations in the Pokemon world
- `explore`: Displays list of Pokemon in map area. Syntax: `explore $AREA_NAME`
- `catch`: Attempt to catch a Pokemon. Syntax: `catch $POKEMON_NAME`
- `inspect`: Show info of pokemon registered in Pokedex. Syntax: `inspect $POKEMON_NAME`
- `pokedex`: Displays all pokemon registered to the Pokedex
- `exit`: Exit the Pokedex

## TO DO:

- [ ] Glamorize the cli to improve readability: colors and back grounds to help lead the eye. (Looking ingto usings the `charmbracelet/lipgloss` package)
- [ ] Organize the `help` command: make the output sorted and return the same order of commands every time.
- [ ] Listen for keyboard arrow inputs: make pressing up and down cycle through command history.
- [ ] More interactive exploration: turn the map command into a selection form that chains into the explore command.
