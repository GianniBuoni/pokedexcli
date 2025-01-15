package api

import (
	"fmt"
)

func CommandMap(c *Config) error {
	if c.Next != nil && *c.Next == "" {
		fmt.Println("you're on the last page")
		fmt.Println()
		return nil
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

func CommandMapb(c *Config) error {
	if c.Previous == nil || *c.Previous == "" {
		fmt.Println("you're on the first page")
		fmt.Println()
		return nil
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
