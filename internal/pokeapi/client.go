package pokeapi

import (
	"net/http"
	"time"

	"github.com/HeHHeyboi/pokedexcli/internal/pokecache"
)

type Client struct {
	Cache  pokecache.Cache
	client http.Client
}

func NewClient(interval, timeout time.Duration) Client {
	new := Client{
		Cache: pokecache.NewCache(interval),
		client: http.Client{
			Timeout: timeout,
		},
	}
	return new

}
