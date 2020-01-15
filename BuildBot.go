package main

import (
	"log"
	"time"

	"github.com/alaingilbert/ogame"
)

// BuildBot Take care of the Build Queues
func BuildBot(bot *ogame.OGame, data *BotData) {
	for i := 0; i < len(data.Planets); i++ {
		for n := 0; i < len(data.Planets[i].BuildQueue); n++ {
			// If the finish Timstamp has been reached and if htere is something in Progress
			if data.Planets[i].Buildinginprogress.FinishAt < time.Now().Unix() && data.Planets[i].Buildinginprogress.FinishAt > 0 {
				log.Println("Building is finished " + data.Planets[i].Buildinginprogress.ID.String())
				data.Planets[i].Buildinginprogress.FinishAt = 0
				var empty ogame.ID
				data.Planets[i].Buildinginprogress.ID = empty
			}

			// If no building is in progress
			if data.Planets[i].Buildinginprogress.FinishAt == 0 {
				err := bot.BuildBuilding(data.Planets[i].ID.Celestial(), data.Planets[i].BuildQueue[n].ID)
				if err == nil {
					log.Println("Start Build on Planet " + data.Planets[i].Name + " " + data.Planets[i].BuildQueue[n].ID.String())
					data.Planets[i].BuildQueue = data.Planets[i].BuildQueue[1:] // Delete first element in Queue
					building, buildingremaining, researchid, researchremaining := bot.ConstructionsBeingBuilt(data.Planets[i].ID.Celestial())
					log.Println("Rsearch: " + researchid.String())
					data.Planets[i].Buildinginprogress.ID = building
					data.Planets[i].Buildinginprogress.FinishAt = GetFinishAt(buildingremaining)
					data.Planets[i].Researchinprogress.ID = researchid
					data.Planets[i].Researchinprogress.FinishAt = GetFinishAt(researchremaining)

				}
			}
		}
	}
}

// GetFinishAt returns the end  timestamp of remaining timer.
func GetFinishAt(t int64) int64 {
	return time.Now().Unix() + t
}
