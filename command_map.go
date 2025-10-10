package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocAreaUrl)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s \n", area.Name)

	}
	cfg.nextLocAreaUrl = resp.Next
	cfg.prevLocAreaUrl = resp.Previous
	return nil
}
func callbackMapB(cfg *config) error {
	if cfg.prevLocAreaUrl == nil {
		return errors.New("you are on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocAreaUrl)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s \n", area.Name)

	}
	cfg.nextLocAreaUrl = resp.Next
	cfg.prevLocAreaUrl = resp.Previous
	return nil
}
