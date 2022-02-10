package main

import (
	//"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/tidwall/gjson"
)

func TestErrorCreateUnit(t *testing.T) {
	var planet = PlanetModel{}
	_, err := createPlanet(planet)
	if err.Error() != "Cannot create empty planet" {
		t.Error("Expected to not create")
	}
}

func TestErrorDeleteUnit(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	_, err := deletePlanetById(strconv.Itoa(rand.Int()))
	if err.Error() != "Planet not found" {
		t.Error("Expected to not delete")
	}
}

func TestErrorSearchUnit(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	_, err := searchPlanetByName(strconv.Itoa(rand.Int()))
	if err.Error() != "No planet found" {
		t.Error("Expected to not to find")
	}
}

func TestErrorGetByIdUnit(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	v := rand.Int()
	_, err := getPlanetById(strconv.Itoa(v))
	if err.Error() != fmt.Sprintf("Id %d not found", v) {
		t.Error("Expected to not to find")
	}
}

func TestCreateDeleteUnit(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
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

func TestGeneralErrorService(t *testing.T) {

	url := fmt.Sprintf("http://%s/api/planetz", ApiEndpoint)

	response, err := http.Get(url)
	if err != nil {
		t.Error("Expected no error. Check if server is up")
	}

	if response.StatusCode != http.StatusNotFound {
		t.Error("Expected 404 not found")
	}
}

func TestOptionsMethodErrorService(t *testing.T) {

	url := fmt.Sprintf("http://%s/api/planets", ApiEndpoint)

	req, err := http.NewRequest("OPTIONS", url, nil)
	if err != nil {
		t.Error("Expected no error. Check if server is up")
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error("Expected no error")
	}

	if err != nil {
		t.Error("Expected no error")
	}

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected %d method not allowed", http.StatusMethodNotAllowed)
	}
}

func TestPutMethodErrorService(t *testing.T) {

	url := fmt.Sprintf("http://%s/api/planets", ApiEndpoint)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		t.Error("Expected no error. Check if server is up")
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error("Expected no error")
	}

	if err != nil {
		t.Error("Expected no error")
	}

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected %d method not allowed", http.StatusMethodNotAllowed)
	}
}

func TestHeadMethodErrorService(t *testing.T) {

	url := fmt.Sprintf("http://%s/api/planets", ApiEndpoint)

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		t.Error("Expected no error. Check if server is up")
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error("Expected no error")
	}

	if err != nil {
		t.Error("Expected no error")
	}

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected %d method not allowed", http.StatusMethodNotAllowed)
	}
}

func TestErrorDeleteService(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	v := rand.Int()

	url := fmt.Sprintf("http://%s/api/planets/%s", ApiEndpoint, strconv.Itoa(v))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Error("Expected no error")
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error("Expected no error")
	}

	if err != nil {
		t.Error("Expected no error")
	}

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected %d not found", http.StatusNotFound)
	}
}

func TestErrorGetService(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	v := rand.Int()

	url := fmt.Sprintf("http://%s/api/planets/%s", ApiEndpoint, strconv.Itoa(v))

	response, err := http.Get(url)
	if err != nil {
		t.Error("Expected no error")
	}

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected %d not found", http.StatusNotFound)
	}
}

func TestErrorSearchService(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	v := rand.Int()

	url := fmt.Sprintf("http://%s/api/planets?name=%s", ApiEndpoint, strconv.Itoa(v))

	response, err := http.Get(url)
	if err != nil {
		t.Error("Expected no error")
	}

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected %d not found", http.StatusNotFound)
	}
}

func TestCreateDeleteService(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	v := rand.Int()

	// Create
	var planet = PlanetModel{
		"",
		fmt.Sprintf("name_%d", v),
		fmt.Sprintf("climate_%d", v),
		fmt.Sprintf("terrain_%d", v),
		int64(v),
	}

	endpoint := fmt.Sprintf("http://%s/api/planets", ApiEndpoint)

	params := make(url.Values)
	params.Set("name", planet.Name)
	params.Set("climate", planet.Climate)
	params.Set("terrain", planet.Terrain)

	response, err := http.PostForm(endpoint, params)

	if err != nil {
		t.Error("Expected no error")
	}

	if response.StatusCode != http.StatusCreated {
		t.Errorf("Expected %d created", http.StatusCreated)
	}

	body, err := io.ReadAll(response.Body)

	createdPlanet := PlanetModel{
		gjson.Get(string(body), "id").String(),
		gjson.Get(string(body), "name").String(),
		gjson.Get(string(body), "climate").String(),
		gjson.Get(string(body), "terrain").String(),
		gjson.Get(string(body), "apparitions").Int(),
	}

	if createdPlanet.Name != planet.Name ||
		createdPlanet.Climate != planet.Climate ||
		createdPlanet.Terrain != planet.Terrain ||
		createdPlanet.Apparitions != 0 { // Custom planet is not present in any Star wars movie
		t.Error("Create service expected to be equal")
	}

	// Get
	endpoint = fmt.Sprintf("http://%s/api/planets/%s", ApiEndpoint, createdPlanet.Id)

	response, err = http.Get(endpoint)
	if err != nil {
		t.Error("Expected no error")
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		t.Error("Expected no error")
	}

	readPlanet := PlanetModel{
		gjson.Get(string(body), "id").String(),
		gjson.Get(string(body), "name").String(),
		gjson.Get(string(body), "climate").String(),
		gjson.Get(string(body), "terrain").String(),
		gjson.Get(string(body), "apparitions").Int(),
	}

	if readPlanet.Id != createdPlanet.Id ||
		readPlanet.Name != planet.Name ||
		readPlanet.Climate != planet.Climate ||
		readPlanet.Terrain != planet.Terrain ||
		readPlanet.Apparitions != 0 { // Custom planet is not present in any Star wars movie
		t.Error("Read service expected to be equal")
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected %d Ok", http.StatusOK)
	}

	// Delete
	endpoint = fmt.Sprintf("http://%s/api/planets/%s", ApiEndpoint, createdPlanet.Id)
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		t.Error("Expected no error")
	}

	response, err = client.Do(req)
	if err != nil {
		t.Error("Expected no error")
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		t.Error("Expected no error")
	}

	status := gjson.Get(string(body), "status").Bool()

	if status != true {
		t.Error("Expected delete to work")
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected %d Ok", http.StatusOK)
	}

	// Get
	endpoint = fmt.Sprintf("http://%s/api/planets/%s", ApiEndpoint, createdPlanet.Id)

	response, err = http.Get(endpoint)
	if err != nil {
		t.Error("Expected no error")
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		t.Error("Expected no error")
	}

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected %d not found", http.StatusNotFound)
	}

	defer response.Body.Close()
}
