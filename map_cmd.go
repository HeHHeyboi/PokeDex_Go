package main

import (
	"fmt"
)

func commandMap(c *Config) error {
	area, err := c.pokeclient.ListLocation(&c.next)
	if err != nil {
		return err
	}
	c.next = area.Next
	c.prev = area.Previous
	for _, v := range area.Results {
		fmt.Println(v.Name)
	}
	return nil

}

func commandMapb(c *Config) error {
	if c.prev == "" {
		return fmt.Errorf("This is First page, Cannot go back anymore.")
	}
	area, err := c.pokeclient.ListLocation(&c.prev)
	if err != nil {
		return err
	}
	c.next = area.Next
	c.prev = area.Previous
	for _, v := range area.Results {
		fmt.Println(v.Name)
	}
	return nil
}
