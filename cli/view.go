package cli

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/bitrise-io/go-utils/colorstring"
	"github.com/bitrise-tools/garden/config"
	"github.com/codegangsta/cli"
)

func printGardenMapView(gardenMap config.GardenMapModel) {
	fmt.Println()
	log.Println("=== Plants ===")
	for plantID, plantModel := range gardenMap.Plants {
		log.Println("🌱 ", colorstring.Green(plantID))
		log.Println("   path:", plantModel.Path)
		log.Println("   seed:", plantModel.Seed)
		log.Println("   vars:", plantModel.Vars)
		log.Println("   zones:", plantModel.Zones)
	}
	log.Println("==============")
}

func view(c *cli.Context) {
	// make it clear what we're viewing
	viewingWhat := colorstring.Green("the whole garden") + "."
	if WorkWithPlantID != "" {
		viewingWhat = colorstring.Yellow("plant") + ": " + WorkWithPlantID + "."
	} else if WorkWithZone != "" {
		viewingWhat = colorstring.Blue("zone") + ": " + WorkWithZone + "."
	}
	log.Infoln("Viewing", viewingWhat)

	gardenMap, err := loadGardenMap()
	if err != nil {
		log.Fatalf("Failed to load Garden Map: %s", err)
	}

	printGardenMapView(gardenMap)
}