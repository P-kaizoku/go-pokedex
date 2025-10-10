package pokeapi

import (
	"net/http"
	"time"

	"github.com/P-kaizoku/go-pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	cache      pokecache.Cache
	HttpClient http.Client
}

func NewClient(interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		HttpClient: http.Client{
			Timeout: time.Minute,
		}}
}
