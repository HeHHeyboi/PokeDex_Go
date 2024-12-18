package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) Explore_location(name string) (*LocationDetail, error) {
	url := baseURL + "/location-area/" + name
	if v, ok := c.Cache.Get(url); ok {
		detail := LocationDetail{}
		if err := json.Unmarshal(v, &detail); err != nil {
			return &LocationDetail{}, err
		}
		return &detail, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return &LocationDetail{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &LocationDetail{}, err
	}

	c.Cache.Add(url, data)
	area_info := &LocationDetail{}
	if err = json.Unmarshal(data, &area_info); err != nil {
		return &LocationDetail{}, err
	}
	return area_info, nil

}
