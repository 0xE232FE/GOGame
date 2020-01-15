package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alaingilbert/ogame"

	"github.com/labstack/echo"
)

// WebBuildBuilding Building
func WebBuildBuilding(c echo.Context) error {
	//	bot := c.Get("bot").(*ogame.OGame)
	data := c.Get("botdata").(*BotData)
	log.Println(&data)
	planetID, err := strconv.ParseInt(c.Param("planetID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ogame.ErrorResp(400, "invalid planet id"))
	}
	ogameID, err := strconv.ParseInt(c.Param("ogameID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ogame.ErrorResp(400, "invalid ogame id"))
	}

	for i := 0; i < len(data.Planets); i++ {
		if data.Planets[i].ID == ogame.PlanetID(planetID) {
			data.Planets[i].AddBuildingtoQueue(ogame.ID(ogameID))
			log.Println("Add to queue " + string(ogameID))
		}
	}
	return c.JSON(http.StatusOK, ogame.SuccessResp(nil))
}
