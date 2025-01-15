package api

import "fmt"

func CommandMap(c *Config) error {
  fmt.Println(c.Next)
  // make GET request
  return nil
}
