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
	Callback    func(c *Config) error
}

type Config struct {
	Next     *string
	Previous *string
	Client   Client
}

type PokeResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []Results
}

type Results struct {
	Name string
	Url  string
}
