package api

type Endpoint = string

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config) error
}

type Config struct {
	Next     Endpoint
	Previous Endpoint
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
