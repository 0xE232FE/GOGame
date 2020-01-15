package main

import "github.com/alaingilbert/ogame"

// ExtMoon is the Version that holds additional information about Moon
type ExtMoon struct {
	ogame.Moon
	ogame.ResourcesBuildings
	ogame.Facilities
	ogame.ShipsInfos
	ogame.DefensesInfos
	Lastupdate int64
}
