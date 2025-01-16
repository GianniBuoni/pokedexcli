package api

import (
	"fmt"
	"os"
)

func commandExit(c *Config, i string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config, i string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}

	fmt.Println()
	return nil
}

func commandPokedex(c *Config, i string) error {
  fmt.Println("Your Pokedex:")
  for _, v := range c.Pokedex {
    fmt.Printf("  - %s\n", v.Name)
  }
	fmt.Println()
  return nil
}
