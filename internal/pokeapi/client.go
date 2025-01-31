package pokeapi

import (
	"net/http"
	"time"

	"github.com/volcente/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheDuration time.Duration) Client {
	return Client{
		cache: *pokecache.CreateCache(cacheDuration),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
