package main

import (
	"log"

	"github.com/alaingilbert/ogame"
)

// Fleetsaver is used to save fleets and control the sending and returning fleets
func Fleetsaver(data *BotData, bot *ogame.OGame) error {
	var planet4saving map[ogame.PlanetID]ogame.Planet = make(map[ogame.PlanetID]ogame.Planet)
	for _, p := range bot.GetCachedPlanets() {
		planet4saving[p.ID] = p
	}

	for _, p := range bot.GetCachedPlanets() {
		for _, f := range data.Fleetsaver {
			if p.Coordinate == f.Origin && f.isFleetsaver {
				delete(planet4saving, p.ID)
			}
		}
	}

	log.Printf("Saving for %d Planets", len(planet4saving))

	return nil
}

func test() {

}
