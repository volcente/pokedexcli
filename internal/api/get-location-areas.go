package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const GetLocationAreasEndpointUrl = "https://pokeapi.co/api/v2/location-area"

type responseBody struct {
	Count         int            `json:"count"`
	NextUrl       *string        `json:"next"`
	PreviousUrl   *string        `json:"previous"`
	LocationAreas []locationArea `json:"results"`
}

type locationArea struct {
	Name string
	Url  string
}

func GetLocationAreas(endpointUrl string) (*responseBody, error) {
	res, err := http.Get(endpointUrl)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch location areas: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}

	var responseBody responseBody
	if err = json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		return nil, fmt.Errorf("Failed to parse response body: %w", err)
	}

	return &responseBody, nil
}
