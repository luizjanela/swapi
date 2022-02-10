package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestErrorCreateUnit(t *testing.T) {
	var planet = PlanetModel{}
	_, err := createPlanet(planet)
	if err.Error() != "Cannot create empty planet" {
		t.Error("Expected to not create")
	}
}

func TestErrorDeleteUnit(t *testing.T) {
	_, err := deletePlanetById(strconv.Itoa(rand.Int()))
	if err.Error() != "Planet not found" {
		t.Error("Expected to not delete")
	}
}

func TestErrorSearchUnit(t *testing.T) {
	_, err := searchPlanetByName(strconv.Itoa(rand.Int()))
	if err.Error() != "No planet found" {
		t.Error("Expected to not to find")
	}
}

func TestErrorGetByIdUnit(t *testing.T) {
	v := rand.Int()
	_, err := getPlanetById(strconv.Itoa(v))
	if err.Error() != fmt.Sprintf("Id %d not found", v) {
		t.Error("Expected to not to find")
	}
}

func TestCreateDeleteUnit(t *testing.T) {
	v := rand.Int()

	var planet = PlanetModel{
		"",
		fmt.Sprintf("name_%d", v),
		fmt.Sprintf("climate_%d", v),
		fmt.Sprintf("terrain_%d", v),
		int64(v),
	}

	createdPlanet, err := createPlanet(planet)

	if createdPlanet.Id == planet.Id {
		t.Error("Expected id to be new")
	}

	if createdPlanet.Name != planet.Name ||
		createdPlanet.Climate != planet.Climate ||
		createdPlanet.Terrain != planet.Terrain ||
		createdPlanet.Apparitions != planet.Apparitions {
		t.Error("Create expected to be equal")
	}

	readPlanet, err := getPlanetById(createdPlanet.Id)

	if createdPlanet.Id != readPlanet.Id ||
		createdPlanet.Name != readPlanet.Name ||
		createdPlanet.Climate != readPlanet.Climate ||
		createdPlanet.Terrain != readPlanet.Terrain ||
		createdPlanet.Apparitions != readPlanet.Apparitions {
		t.Error("Read expected to be equal")
	}

	searchPlanet, err := getPlanetById(createdPlanet.Id)

	if createdPlanet.Id != searchPlanet.Id ||
		createdPlanet.Name != searchPlanet.Name ||
		createdPlanet.Climate != searchPlanet.Climate ||
		createdPlanet.Terrain != searchPlanet.Terrain ||
		createdPlanet.Apparitions != searchPlanet.Apparitions {
		t.Error("Search expected to be equal")
	}

	result, err := deletePlanetById(createdPlanet.Id)
	if result == false {
		t.Error("Expected to delete")
	}

	_, err = getPlanetById(createdPlanet.Id)
	if err != nil && err.Error() != fmt.Sprintf("Id %s not found", createdPlanet.Id) {
		t.Error("Expected to not find")
	}

	_, err = searchPlanetByName(createdPlanet.Name)
	if err != nil && err.Error() != "No planet found" {
		t.Error("Expected to not to find")
	}
}
