package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c Client) GetLocationPokemon(locationName string) (locationPokemonResp, error) {
	url := fmt.Sprintf("%s%s/%s", baseURL, locationAreasEndpoint, locationName)

	if cachedData, wasFound := c.cache.Get(url); wasFound {
		var data locationPokemonResp
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return locationPokemonResp{}, fmt.Errorf("Failed to parse cached data: %w", err)
		}
		return data, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return locationPokemonResp{}, fmt.Errorf("Failed to fetch location pokemon: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return locationPokemonResp{}, fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationPokemonResp{}, fmt.Errorf("Failed to read response body: %w", err)
	}

	var data locationPokemonResp
	if err = json.Unmarshal(body, &data); err != nil {
		return locationPokemonResp{}, fmt.Errorf("Failed to parse response body: %w", err)
	}
	c.cache.Add(url, body)

	return data, nil
}

type locationPokemonResp struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
