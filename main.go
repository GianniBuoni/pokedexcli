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
	prompt()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatalf("There was trouble reading input: %v", err)
		}

		input, err := cleanInput(scanner.Text())
		if err != nil {
			fmt.Print(err)
			prompt()
			continue
		}

		s := fmt.Sprintf("Your command was: %s", input[0])
		fmt.Println(s)
		prompt()
	}
}
