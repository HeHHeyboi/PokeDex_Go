package main

import "fmt"

var cmdOrder = []string{"help", "exit", "map", "mapb", "explore", "catch", "inspect", "pokedex"}

func commandHelp(c *Config) error {
	fmt.Println("Welcomd to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, n := range cmdOrder {
		command := cmd[n]
		fmt.Printf("%s: %s\n", command.name, command.description)

	}
	fmt.Println()
	return nil
}
