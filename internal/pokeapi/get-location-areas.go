package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetLocationAreas(endpointUrl *string) (locationAreasResp, error) {
	url := fmt.Sprintf("%s%s", baseURL, locationAreasEndpoint)
	if endpointUrl != nil {
		url = *endpointUrl
	}

	if cachedData, wasFound := c.cache.Get(url); wasFound {
		var data locationAreasResp
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return locationAreasResp{}, fmt.Errorf("Failed to parse cached data: %w", err)
		}
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return locationAreasResp{}, fmt.Errorf("Failed to fetch location areas: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return locationAreasResp{}, fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreasResp{}, fmt.Errorf("Failed to read response body: %w", err)
	}

	var data locationAreasResp
	if err = json.Unmarshal(body, &data); err != nil {
		return locationAreasResp{}, fmt.Errorf("Failed to parse response body: %w", err)
	}
	c.cache.Add(url, body)

	return data, nil
}

type locationAreasResp struct {
	Count         int     `json:"count"`
	NextUrl       *string `json:"next"`
	PreviousUrl   *string `json:"previous"`
	LocationAreas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
