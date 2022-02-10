package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

func getApparitionsByName(name string) (int64, error) {

	/*
		Direct SWAPI Request
		https://swapi.dev/api/planets/?search=Tatooine
	*/

	url := fmt.Sprintf("%s/api/planets/?search=%s", SwHost, name)

	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	count := gjson.Get(string(body), "count").Int()

	if count == 0 {
		return 0, fmt.Errorf("Planet %s does not have any apparitions in Star wars", name)
	}

	apparitions := gjson.Get(string(body), "results.0.films.#").Int()

	defer response.Body.Close()

	return apparitions, nil
}
