package repl

import (
	"fmt"
)

func mapBackCommand(config *Config) error {
	if config.prevLocationsURL == nil {
		fmt.Println("You're on the first page.")
		return nil
	}

	responseBody, err := config.pokeClient.GetLocationAreas(config.prevLocationsURL)
	if err != nil {
		return err
	}

	for _, location := range responseBody.LocationAreas {
		fmt.Println(location.Name)
	}

	config.prevLocationsURL = responseBody.PreviousUrl
	config.nextLocationsURL = responseBody.NextUrl

	return nil
}
