package main

import "fmt"

func commandHelp(command map[string]cliCommand) error {
	fmt.Println("Welcomd to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, v := range command {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	fmt.Println()
	return nil
}
