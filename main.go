package main

import (
	"time"

	"github.com/HeHHeyboi/pokedexcli/internal/pokeapi"
)

var cmd map[string]cliCommand

func init() {
	cmd = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: func(c *Config, arg string) error {
				return commandHelp(c)
			},
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback: func(c *Config, arg string) error {
				return commandExit(c)
			},
		},
		"map": {
			name:        "map",
			description: "show next 20 location areas in the Pokemon world",
			callback: func(c *Config, arg string) error {
				return commandMap(c)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "show previous 20 location areas in the Pokemon world",
			callback: func(c *Config, arg string) error {
				return commandMapb(c)
			},
		},
		"explore": {
			name: "explore",
			description: `explore <name> list pokemon that live inside location
	<name> : name of the city or location`,
			callback: func(c *Config, name string) error {
				return commandExplore(c, name)
			},
		},
		"catch": {
			name: "catch",
			description: `catch a pokemon with given name
	command: catch <name>
	name : pokemon's name`,
			callback: func(c *Config, name string) error {
				return commandCatch(c, name)
			},
		},
		"inspect": {
			name: "inspect",
			description: `show info about pokemon
	inspect <name>
	name: pokemon's name`,
			callback: func(c *Config, name string) error {
				return commandInspect(c, name)
			},
		},
	}
}

func main() {
	client := pokeapi.NewClient(10*time.Second, time.Minute)
	cfg := Config{
		pokeclient: client,
	}
	startRepl(&cfg)
}
