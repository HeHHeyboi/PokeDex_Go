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

var cmd map[string]cliCommand

func init() {
	cmd = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: func(c *Config, arg string) error {
				return commandHelp(cmd, c)
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
	}
}
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
