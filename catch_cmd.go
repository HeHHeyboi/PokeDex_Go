package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO: add Chance to catch pokemon using "math/rand" & and chance is base or experience
func commandCatch(c *Config, name string) error {
	rand.NewSource(time.Now().UnixNano())
	if name == "" {
		return fmt.Errorf("Please input name of pokemon")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := c.pokeclient.Get_Pokemon(name)
	if err != nil {
		return err
	}
	chance := 50 - int(pokemon.BaseExperience/10)
	if rand.Intn(100) < chance {
		fmt.Println(pokemon.Name, " was caught!")
		fmt.Println("You may now inspect it with 'inspect' command.")
		pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name, " was escape!")
	}

	return nil
}
