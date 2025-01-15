package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	. "github.com/GianniBuoni/pokedexcli/internal/api"
	. "github.com/GianniBuoni/pokedexcli/internal/lib"
	. "github.com/GianniBuoni/pokedexcli/internal/pokecache"
)

const (
	mapArea Endpoint = "https://pokeapi.co/api/v2/location-area?limit=20"
)

func prompt() {
	fmt.Print("Pokedex > ")
}

func main() {
  NewCache(time.Second * 5)
	var endpoints Config
  endpoints.Next = mapArea
	scanner := bufio.NewScanner(os.Stdin)

	for {
		prompt()
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatalf("There was trouble reading input: %v", err)
		}

		input, err := CleanInput(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}

		command, ok := GetCommands()[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		command.Callback(&endpoints)
		continue
	}
}
