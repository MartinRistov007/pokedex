package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (LocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationArea{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationResp := LocationArea{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, data)

	return locationResp, nil
}
