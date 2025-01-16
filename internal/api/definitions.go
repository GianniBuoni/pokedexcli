package api

import (
	cache "github.com/GianniBuoni/pokedexcli/internal/pokecache"
)

type Client struct {
	cache *cache.Cache
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config, i string) error
}

type Config struct {
	Next     *string
	Previous *string
	Client   Client
	Pokedex  map[string]Pokemon
}

type LocationResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []Results
}

type PokeLocations struct {
	PokemonEncounters []struct {
		Pokemon Results `mapstructure:"pokemon"`
	} `mapstructure:"pokemon_encounters"`
}

type Pokemon struct {
	Id      int    `mapstructure:"id"`
	BaseExp int    `mapstructure:"base_experience"`
	Name    string `mapstructure:"name"`
}

type Results struct {
	Name string
	Url  string
}
