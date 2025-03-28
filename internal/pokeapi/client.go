package pokeapi

import (
	"net/http"
	"time"

	"github.com/Axelandrovitch/pokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
	baseURL    string
}

func NewClient(baseURL string, timeout, cacheDuration time.Duration) Client {
	cache := pokecache.NewCache(cacheDuration)
	return Client{
		cache:      *cache,
		httpClient: http.Client{Timeout: timeout},
		baseURL:    baseURL,
	}
}
