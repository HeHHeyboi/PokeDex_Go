package main

import "fmt"

func commandInspect(c *Config, name string) error {
	pokemon, ok := pokedex[name]
	if !ok {
		return fmt.Errorf("you haven't caught that pokemon")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)

	fmt.Println("Status:")
	for _, v := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", v.Stat.Name, v.BaseStat)
	}

	fmt.Println("Types:")
	for _, v := range pokemon.Types {
		fmt.Printf("  - %s\n", v.Type.Name)
	}

	return nil
}
