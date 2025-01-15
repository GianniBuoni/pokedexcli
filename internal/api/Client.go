package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	. "github.com/GianniBuoni/pokedexcli/internal/pokecache"
)

const (
	baseUrl string = "https://pokeapi.co/api/v2/"
)

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: NewCache(cacheInterval),
	}
}

func (c *Client) getLocation(path *string) (PokeResponse, error) {
	fullpath := baseUrl + "location-area?limit=20"
	if path != nil {
		fullpath = *path
	}

	data, err := c.request(fullpath)
	if err != nil {
		return PokeResponse{}, err
	}
	return data, nil
}

func (c *Client) request(url string) (PokeResponse, error) {
	// check cache
	if val, ok := c.cache.Get(url); ok {
		data := PokeResponse{}
		err := json.Unmarshal(val, &data)
		if err != nil {
			return PokeResponse{}, fmt.Errorf("issue reading cache: %w", err)
		}
		return data, nil
	}

	// make GET request
	res, err := http.Get(url)
	if err != nil {
		return PokeResponse{}, fmt.Errorf("issue with GET request: %w", err)
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return PokeResponse{}, fmt.Errorf(
			"issue reading response\nstatus code: %d\nbody: %v",
			res.StatusCode, res.Body)
	}

	data := PokeResponse{}
	bytes, err := io.ReadAll(res.Body)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return PokeResponse{}, err
	}

	c.cache.Add(url, bytes)
	return data, nil
}
