package api

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func commandExplore(c *Config, i string) error {
	if i == "" {
		return fmt.Errorf("area param cannot be empty")
	}

	// get pokemon data
	fmt.Printf("Exploring %s...\n", i)
	data, err := c.Client.getPokemonInArea(i)
	if err != nil {
		return err
	}

	// list pokemon
	for _, pokemon := range data.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	fmt.Println()

	return nil
}

func (c *Client) getPokemonInArea(name string) (PokeLocations, error) {
	fullPath := baseUrl + "location-area/" + name

	data, err := c.request(fullPath)
	if err != nil {
		return PokeLocations{}, err
	}

	pokeData := PokeLocations{}
	err = mapstructure.Decode(data, &pokeData)
	if err != nil {
		return PokeLocations{}, fmt.Errorf("issue decoding map to structure: %w", err)
	}

	return pokeData, nil
}
