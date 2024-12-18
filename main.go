package main

import (
	"time"

	"github.com/HeHHeyboi/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(10*time.Second, time.Minute)
	cfg := Config{
		pokeclient: client,
	}

	startRepl(&cfg)
}
