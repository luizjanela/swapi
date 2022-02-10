package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

func createPlanet(planet PlanetModel) (PlanetModel, error) {

	/*
		Direct ELK Request
		curl -X POST http://127.0.0.1:9200/index-swapi-0001/_doc --data '{"name": "Terra nova", "climate": "temperate", "terrain": "grasslands, mountains"}' -H 'Content-Type: application/json'
	*/

	url := fmt.Sprintf("%s/_doc", ElkHost)

	// Todo exclude id from empty Marshal
	postBody, err := json.Marshal(planet)

	response, err := http.Post(url, "application/json", strings.NewReader(string(postBody)))
	if err != nil {
		return PlanetModel{}, err
	}

	body, err := io.ReadAll(response.Body)

	result := gjson.Get(string(body), "result").String()

	if result != "created" {
		return PlanetModel{}, fmt.Errorf("Error creating planet")
	}

	planet.Id = gjson.Get(string(body), "_id").String()

	defer response.Body.Close()

	return planet, nil
}

func getPlanetById(id string) (PlanetModel, error) {

	/*
		Direct ELK Request
		curl http://127.0.0.1:9200/index-swapi-0001/_doc/8gS22n4BYh62MkvteAsR1?pretty
	*/

	var planet PlanetModel

	url := fmt.Sprintf("%s/_doc/%s", ElkHost, id)

	response, err := http.Get(url)
	if err != nil {
		return planet, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return planet, err
	}

	found := gjson.Get(string(body), "found").Bool()

	if found == false {
		return planet, fmt.Errorf("Id %s not found", id)
	}

	planet = PlanetModel{
		gjson.Get(string(body), "_id").String(),
		gjson.Get(string(body), "_source.name").String(),
		gjson.Get(string(body), "_source.climate").String(),
		gjson.Get(string(body), "_source.terrain").String(),
		gjson.Get(string(body), "_source.apparitions").Int(),
	}

	defer response.Body.Close()

	return planet, nil
}

// Todo implement limit offset logic
func listPlanets() ([]PlanetModel, error) {

	/*
		Direct ELK Request
		curl http://127.0.0.1:9200/index-swapi-0001/_search?pretty
	*/

	var planets []PlanetModel

	// Todo implement limit offset logic
	postBody := `{"size":100}`

	url := fmt.Sprintf("%s/_search", ElkHost)

	response, err := http.Post(url, "application/json", strings.NewReader(postBody))
	if err != nil {
		return planets, err
	}

	body, err := io.ReadAll(response.Body)

	total := gjson.Get(string(body), "hits.total.value").Int()

	if total == 0 {
		return planets, nil
	}

	var planet PlanetModel

	result := gjson.Get(string(body), "hits.hits")
	result.ForEach(func(key, value gjson.Result) bool {

		planet = PlanetModel{
			gjson.Get(value.String(), "_id").String(),
			gjson.Get(value.String(), "_source.name").String(),
			gjson.Get(value.String(), "_source.climate").String(),
			gjson.Get(value.String(), "_source.terrain").String(),
			gjson.Get(value.String(), "_source.apparitions").Int(),
		}

		planets = append(planets, planet)

		return true
	})

	defer response.Body.Close()

	return planets, nil
}

func searchPlanets(name string) ([]PlanetModel, error) {

	/*
		Direct ELK Request
		curl http://127.0.0.1:9200/index-swapi-0001/_search?pretty --data '{"query":{"match":{"name": "Alderaan"}}}' -H 'Content-Type: application/json'
	*/

	var planets []PlanetModel

	postBody := fmt.Sprintf(`{"query":{"match":{"name": "%s"}}}`, name)

	url := fmt.Sprintf("%s/_search", ElkHost)

	response, err := http.Post(url, "application/json", strings.NewReader(postBody))
	if err != nil {
		return planets, err
	}

	body, err := io.ReadAll(response.Body)

	total := gjson.Get(string(body), "hits.total.value").Int()

	if total == 0 {
		return planets, fmt.Errorf("No planets found")
	}

	var planet PlanetModel

	result := gjson.Get(string(body), "hits.hits")
	result.ForEach(func(key, value gjson.Result) bool {

		planet = PlanetModel{
			gjson.Get(value.String(), "_id").String(),
			gjson.Get(value.String(), "_source.name").String(),
			gjson.Get(value.String(), "_source.climate").String(),
			gjson.Get(value.String(), "_source.terrain").String(),
			gjson.Get(value.String(), "_source.apparitions").Int(),
		}

		planets = append(planets, planet)

		return true
	})

	defer response.Body.Close()

	return planets, nil
}

func deletePlanetById(id string) (bool, error) {

	/*
		Direct ELK Request
		curl http://127.0.0.1:9200/index-swapi-0001/_doc/8gS22n4BYh62MkvteAsR1
	*/

	url := fmt.Sprintf("%s/_doc/%s", ElkHost, id)

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return false, err
	}

	response, err := client.Do(req)
	if err != nil {
		return false, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	result := gjson.Get(string(body), "result").String()

	if result != "deleted" {
		return false, fmt.Errorf("Id %s could not be deleted", id)
	}

	defer response.Body.Close()

	return true, nil
}
