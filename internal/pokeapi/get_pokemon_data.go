package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) Get_Pokemon(name string) (*Pokemon_Info, error) {
	url := pokemonURL + name
	// TODO: Add Cache for pokemonURL & and check if pokemonURL is in the Cache or not

	res, err := http.Get(url)
	if err != nil {
		return &Pokemon_Info{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &Pokemon_Info{}, err
	}

	var poke_info Pokemon_Info
	if err = json.Unmarshal(data, &poke_info); err != nil {
		return &Pokemon_Info{}, err
	}

	return &poke_info, nil
}
