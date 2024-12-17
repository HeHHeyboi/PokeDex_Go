package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())[0]
		fmt.Println("Your command was:", input)
	}
}

func cleanInput(text string) []string {
	lowerString := strings.ToLower(text)
	out := strings.Fields(lowerString)
	return out
}
