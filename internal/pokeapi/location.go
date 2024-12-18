package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (client *Client) ListLocation(pageUrl *string) (Location, error) {
	url := baseURL + "/location-area"
	if *pageUrl != "" {
		url = *pageUrl
	}
	if val, ok := client.Cache.Get(url); ok {
		area := Location{}
		fmt.Println(url)
		err := json.Unmarshal(val, &area)
		if err != nil {
			return Location{}, err
		}
		for _, loc := range area.Results {
			fmt.Println(loc.Name)
		}
		fmt.Println(area.Next)
		fmt.Println(area.Previous)
		return area, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}
	area := Location{}
	data, err := io.ReadAll(res.Body)
	client.Cache.Add(url, data)
	fmt.Println("Area url: ", url)

	err = json.Unmarshal(data, &area)
	fmt.Println("Next area: ", area.Next)
	fmt.Println("Previous area: ", area.Previous)
	if err != nil {
		return Location{}, err
	}
	return area, nil

}
