package commands

import (
	"fmt"

	"github.com/volcente/pokedexcli/internal/api"
)

func mapCommand(mapConfig *MapConfig) error {
	endpointUrl := api.GetLocationAreasEndpointUrl
	if mapConfig.NextUrl != nil {
		endpointUrl = *mapConfig.NextUrl
	}

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
