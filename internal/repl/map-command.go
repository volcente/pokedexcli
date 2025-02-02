package repl

import (
	"fmt"
)

func mapCommand(config *Config, commandArgs ...string) error {
	responseBody, err := config.pokeClient.GetLocationAreas(config.nextLocationsURL)
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
