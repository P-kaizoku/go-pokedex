package main

import (
	"time"

	"github.com/P-kaizoku/go-pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient  pokeapi.Client
	nextLocAreaUrl *string
	prevLocAreaUrl *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	startRepl(&cfg)
}
