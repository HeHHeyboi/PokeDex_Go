package main

import (
	"encoding/json"
	"fmt"
	"github.com/HeHHeyboi/pokedexcli/internal/pokecache"
	"net/http"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *Config) error {
	res, err := http.Get(c.next)
	if err != nil {
		return err
	}
	area := Location{}
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&area); err != nil {
		return err
	}
	for _, loc := range area.Results {
		fmt.Println(loc.Name)
	}
	c.next = area.Next
	c.prev = area.Previous
	fmt.Printf("Next Area: %s\n", c.next)
	fmt.Printf("Previous Area: %s\n", c.prev)
	return nil

}
func commandMapb(c *Config) error {
	if c.prev == "" {
		return fmt.Errorf("This is First page, Cannot go back anymore.")
	}
	res, err := http.Get(c.prev)
	if err != nil {
		return err
	}

	area := Location{}
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&area); err != nil {
		return err
	}
	for _, loc := range area.Results {
		fmt.Println(loc.Name)
	}

	c.next = area.Next
	c.prev = area.Previous
	return nil
}
