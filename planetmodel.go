package main

import (
	"fmt"
)

type PlanetModel struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Climate     string `json:"climate"`
	Terrain     string `json:"terrain"`
	Apparitions int64  `json:"apparitions"`
}

func planetModelMarshall(planet PlanetModel) string {
	return fmt.Sprintf(`{"name":"%s", "climate":"%s", "terrain":"%s", "apparitions":%d}`, planet.Name, planet.Climate, planet.Terrain, planet.Apparitions)
}
