package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResop, error) {
	endpoint := "location-area/"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	dat, ok := c.cache.Get(fullURL)
	if ok {
		locationAreaResp := LocationAreaResop{}
		err := json.Unmarshal(dat, &locationAreaResp)
		if err != nil {
			return LocationAreaResop{}, nil
		}

		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResop{}, nil
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return LocationAreaResop{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResop{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResop{}, nil
	}

	locationAreaResp := LocationAreaResop{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResop{}, nil
	}

	c.cache.Add(fullURL, data)

	return locationAreaResp, nil

}
