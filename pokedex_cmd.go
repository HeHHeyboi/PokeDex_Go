package main

import "fmt"

func commandPokedex() error {
	fmt.Println("Your Pokedex:")
	if len(pokedex) == 0 {
		return fmt.Errorf("You didn't caught anything yet")
	}

	for k := range pokedex {
		fmt.Println(" -", k)
	}
	return nil
}
