package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Region struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (client *Client) ListLocation(pageUrl *string) (Region, error) {
	url := baseURL + "/location-area"
	if *pageUrl != "" {
		url = *pageUrl
	}
	if val, ok := client.Cache.Get(url); ok {
		area := Region{}
		err := json.Unmarshal(val, &area)
		if err != nil {
			return Region{}, err
		}
		return area, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return Region{}, err
	}
	area := Region{}
	data, err := io.ReadAll(res.Body)
	client.Cache.Add(url, data)

	err = json.Unmarshal(data, &area)
	if err != nil {
		return Region{}, err
	}
	return area, nil

}
