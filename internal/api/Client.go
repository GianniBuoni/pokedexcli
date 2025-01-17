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

func (c *Client) request(url string) (T map[string]any, err error) {
	// check cache
	if val, ok := c.cache.Get(url); ok {
		data := T
		err := json.Unmarshal(val, &data)
		if err != nil {
			return nil, fmt.Errorf("issue reading cache: %w", err)
		}
		return data, nil
	}

	// make GET request
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("issue with GET request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("Ooops! Nothing was found!")
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf(
			"issue reading response\nstatus code: %d\nbody: %v",
			res.StatusCode, res.Body)
	}

	data := T
	bytes, err := io.ReadAll(res.Body)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, bytes)
	return data, nil
}
