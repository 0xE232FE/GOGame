package main

import (
	"log"
	"time"

	"github.com/alaingilbert/ogame"
)

// ExtPlanet is the Version that holds additional information about a Planet
type ExtPlanet struct {
	ogame.Planet
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

//updateLastupdate Updates last Activitiy for Planet
func (e *ExtPlanet) updateLastupdate() {
	e.Lastupdate = time.Now().Unix()
}

// AddBuildingtoQueue adds a Building to Queue
func (e *ExtPlanet) AddBuildingtoQueue(id ogame.ID, nbr int64) {
	log.Println("Add to queue " + id.String())
	e.BuildQueue = append(e.BuildQueue, ogame.Quantifiable{ID: id, Nbr: nbr})
}
