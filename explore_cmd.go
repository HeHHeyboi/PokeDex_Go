package main

import (
	"fmt"
	"time"
)

func commandExplore(c *Config, name string) error {
	if name == "" {
		return fmt.Errorf("Please input city name")
	}
	data, err := c.pokeclient.Explore_location(name)
	if err != nil {
		return err
	}
	fmt.Println("Exploring ", name, "...")
	time.Sleep(1 * time.Second)
	fmt.Println("Found Pokemon:")
	for _, v := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", v.Pokemon.Name)
	}
	return nil
}
