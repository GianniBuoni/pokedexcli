package api

import (
	"fmt"
	"math/rand"

	"github.com/mitchellh/mapstructure"
)

func catch(c *Config, i string) error {
	if i == "" {
		return fmt.Errorf("pokemon param cannot be empty")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", i)

	data, err := c.Client.getPokemon(i)
	if err != nil {
		return err
	}

	res := rand.Intn(340)
	if res >= data.BaseExp {
		c.Pokedex[i] = data
		fmt.Printf("%s was caught!\n", i)
	} else {
		fmt.Printf("%s escaped!\n", i)
	}

	fmt.Println()
	return nil
}

func (c *Client) getPokemon(name string) (Pokemon, error) {
	fullUrl := baseUrl + "pokemon/" + name
	data, err := c.request(fullUrl)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = mapstructure.Decode(data, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("issue decoding map to struct: %w", err)
	}

	return pokemon, nil
}
