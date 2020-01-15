package main

import (
	"time"

	"github.com/alaingilbert/ogame"
)

// BotData holds Data for everything
type BotData struct {
	Planets []ExtPlanet
	Moons   []ExtMoon
	Fleet   []ogame.Fleet
}

// AddPlanets adds Planets to the BotData
func (e *BotData) AddPlanets(planets []ogame.Planet) {
	for _, planet := range planets {
		e.Planets = append(e.Planets, ExtPlanet{
			Planet:     planet,
			Lastupdate: time.Now().Unix(),
		})
	}
}
