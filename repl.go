package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const baseURL = "https://pokeapi.co/api/v2/location-area/"

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}

type Config struct {
	next string
	prev string
}

var cmd map[string]cliCommand

func init() {
	cmd = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: func(c *Config) error {
				return commandHelp(cmd, c)
			},
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback: func(c *Config) error {
				return commandExit(c)
			},
		},
		"map": {
			name:        "map",
			description: "show next 20 location areas in the Pokemon world",
			callback: func(c *Config) error {
				return commandMap(c)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "show previous 20 location areas in the Pokemon world",
			callback: func(c *Config) error {
				return commandMapb(c)
			},
		},
	}
}
func startRepl() {
	config := Config{baseURL, ""}
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
			err := command.callback(&config)
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
