package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/HeHHeyboi/pokedexcli/internal/pokeapi"
)

const baseURL = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config, arg string) error
}

type Config struct {
	pokeclient pokeapi.Client
	next       string
	prev       string
}

var pokedex map[string]*pokeapi.Pokemon_Info

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, ok := cmd[commandName]
		if ok {
			var err error
			if len(words) > 1 {
				err = command.callback(cfg, words[1])
			} else {
				err = command.callback(cfg, "")
			}

			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowerString := strings.ToLower(text)
	out := strings.Fields(lowerString)
	return out
}
