package main

import (
	"time"

	"github.com/alaingilbert/ogame"
)

// BuildBot Take care of the Build Queues
func BuildBot(bot *ogame.OGame, data *BotData) {
	for i, p := range data.Planet {
		res, _ := bot.GetResourcesBuildings(p.ID.Celestial())
		data.Planet[i].ResourcesBuildings = res
	}
}

// GetFinishAt returns the end  timestamp of remaining timer.
func GetFinishAt(t int64) int64 {
	return time.Now().Unix() + t
}
