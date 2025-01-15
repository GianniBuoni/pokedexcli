package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	. "github.com/GianniBuoni/pokedexcli/internal/api"
	. "github.com/GianniBuoni/pokedexcli/internal/lib"
)

func prompt() {
	fmt.Print("Pokedex > ")
}

func main() {
	config := &Config{
		Client: NewClient(1 * time.Minute),
	}
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

		command.Callback(config)
		continue
	}
}
