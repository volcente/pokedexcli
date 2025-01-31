package commands

import (
	"fmt"

	"github.com/volcente/pokedexcli/internal/api"
)

func mapBackCommand(mapConfig *MapConfig) error {
	if mapConfig.PreviousUrl == nil {
		fmt.Println("You're on the first page.")
		return nil
	}
	endpointUrl := *mapConfig.PreviousUrl

	responseBody, err := api.GetLocationAreas(endpointUrl)
	if err != nil {
		return err
	}

	for _, location := range responseBody.LocationAreas {
		fmt.Println(location.Name)
	}

	mapConfig.PreviousUrl = responseBody.PreviousUrl
	mapConfig.NextUrl = responseBody.NextUrl

	return nil
}
