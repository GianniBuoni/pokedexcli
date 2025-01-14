package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func prompt() {
	fmt.Print("Pokedex > ")
}

func cleanInput(text string) ([]string, error) {
	if text == "" {
		return nil, fmt.Errorf("input string cannot be empty")
	}
	return strings.Fields(strings.ToLower(strings.TrimSpace(text))), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		prompt()
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatalf("There was trouble reading input: %v", err)
		}

		input, err := cleanInput(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}

		command, ok := allCommands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		command.callback()
		continue
	}
}
