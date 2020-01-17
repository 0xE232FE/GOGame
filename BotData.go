package main

import (
	"github.com/alaingilbert/ogame"
)

// BotData holds Data for everything
type BotData struct {
	Planet       []ExtPlanet
	Fleets       []ogame.Fleet
	Fleetsaver   []FleetsaverData
	AttackEvents []ogame.AttackEvent
	FleetEvents  []ogame.Fleet
	Researches   ogame.Researches
}

// ExtPlanet is the Version that holds additional information about a Planet
type ExtPlanet struct {
	ogame.Planet
	ogame.Resources
	ogame.ResourcesBuildings
	ogame.Facilities
	ogame.ShipsInfos
	ogame.DefensesInfos
	BuildQueue         []ogame.Quantifiable
	ShipyardQueue      []ogame.Quantifiable
	Buildinginprogress struct {
		ogame.ID
		FinishAt int64
	}
	Researchinprogress struct {
		ogame.ID
		FinishAt int64
	}
	Shipyardinprogress struct {
		ogame.ID
		FinishAt int64
	}
	Lastupdate int64
}

// FleetsaverData is a modification of ogame.Fleet this will hold all Fleets from the Fleetsaver
type FleetsaverData struct {
	ogame.Fleet
	ArriveAt     int64
	DispatchedAt int64
	recallAt     int64
	isFleetsaver bool
}

func (e *ExtPlanet) updateBuildings(bot *ogame.OGame) {
	res, _ := bot.GetResourcesBuildings(e.ID.Celestial())
	fac, _ := bot.GetFacilities(e.ID.Celestial())
	e.ResourcesBuildings = res
	e.Facilities = fac
}
