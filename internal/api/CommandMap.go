package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CommandMap(c *Config) error {
	if c.Next == "" {
		fmt.Println("you're on the last page")
		fmt.Println()
		return nil
	}

	data, err := mapRequest(c.Next, c)
	if err != nil {
		return err
	}

	for _, v := range data.Results {
		fmt.Println(v.Name)
	}
	fmt.Println()
	return nil
}

func CommandMapb(c *Config) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		fmt.Println()
		return nil

	}

	data, err := mapRequest(c.Previous, c)
	if err != nil {
		return err
	}

	for _, v := range data.Results {
		fmt.Println(v.Name)
	}
	fmt.Println()
	return nil
}

func mapRequest(url string, c *Config) (PokeResponse, error) {
	var data PokeResponse

	// make GET request
	res, err := http.Get(url)
	if err != nil {
		return data, fmt.Errorf("issue with GET request: %w", err)
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return data, fmt.Errorf(
			"issue reading response\nstatus code: %d\nbody: %v",
			res.StatusCode, res.Body)
	}

	// decode res into PokeResponse struct
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return data, fmt.Errorf("issue decoding response body: %w", err)
	}

	// mutate c.Previous and c.Next
	c.Next = data.Next
	c.Previous = data.Previous
	return data, nil
}
