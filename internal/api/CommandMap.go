package api

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func CommandMap(c *Config, i string) error {
	if c.Next != nil && *c.Next == "" {
		return fmt.Errorf("you're on the last page")
	}

	data, err := c.Client.getLocation(c.Next)
	if err != nil {
		return err
	}

	c.Previous = &data.Previous
	c.Next = &data.Next

	for _, v := range data.Results {
		fmt.Println(v.Name)
	}
	fmt.Println()
	return nil
}

func CommandMapb(c *Config, i string) error {
	if c.Previous == nil || *c.Previous == "" {
		return fmt.Errorf("you're on the first page")
	}

	data, err := c.Client.getLocation(c.Previous)
	if err != nil {
		return err
	}

	c.Previous = &data.Previous
	c.Next = &data.Next

	for _, v := range data.Results {
		fmt.Println(v.Name)
	}
	fmt.Println()
	return nil
}

func (c *Client) getLocation(path *string) (LocationResponse, error) {
	fullpath := baseUrl + "location-area?limit=20"
	if path != nil {
		fullpath = *path
	}

	data, err := c.request(fullpath)
	if err != nil {
		return LocationResponse{}, err
	}

	mapData := LocationResponse{}
	err = mapstructure.Decode(data, &mapData)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("issue decoding map to structure: %w", err)
	}

	return mapData, nil
}
