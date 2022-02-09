package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"io"
	"encoding/json"
)

const ElkHost = "http://127.0.0.1:9200/index-swapi-0001"

type PlanetModel struct {
	Id 					string `json:"id"`
	Name 				string `json:"name"`
	Climate 			string `json:"climate"`
	Terrain 			string `json:"terrain"`
	FilmApparitions 	string `json:"film_apparitions"`
}

type ElkModel struct {
	Id 		string 		`json:"_id"`
	Index 	string 		`json:"_index"`
	Planet 	PlanetModel	`json:"_source"`
}

func planetCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("planetCreateHandler")
	fmt.Println(r.URL.Path)
	fmt.Println(r.Method)

	name := r.FormValue("name")
	climate := r.FormValue("climate")
	terrain := r.FormValue("terrain")

	fmt.Println(w, "Planet name: %v\n", name)
	fmt.Println(w, "Planet climate: %v\n", climate)
	fmt.Println(w, "Planet terrain: %v\n", terrain)

	// curl -X POST http://127.0.0.1:9200/index-swapi-0001/_doc --data '{"name": "Alderaan", "climate": "temperate", "terrain": "grasslands, mountains"}' -H 'Content-Type: application/json'
}

func planetListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("planetListHandler")
	fmt.Println(r.URL.Path)
	fmt.Println(r.Method)


	vars := mux.Vars(r)
	fmt.Println(w, "Planet name to search: %v\n", vars["name"])

	//curl http://127.0.0.1:9200/index-swapi-0001/_search?pretty --data '{"query":{"term":{"name": { "value": "Tattoine"}}}}' -H 'Content-Type: application/json'
}

func planetReadHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)

	// curl http://127.0.0.1:9200/index-swapi-0001/_doc/8gS22n4BYh62MkvteAsR?pretty
	url := fmt.Sprintf("%s/_doc/%s", ElkHost, vars["id"])
	
	response, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	body, err := io.ReadAll(response.Body)

	var elk ElkModel

	err = json.Unmarshal(body, &elk)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	elk.Planet.Id = elk.Id

	data, err := json.Marshal(elk.Planet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func planetDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("planetDeleteHandler")
	fmt.Println(r.URL.Path)
	fmt.Println(r.Method)

	vars := mux.Vars(r)
	fmt.Println(w, "Planet id to delete: %v\n", vars["id"])
}

func main() {

	fmt.Println("Starting SWAPI Server... May the force be with us!")

	r := mux.NewRouter()

	r.HandleFunc("/api/planets", planetCreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/planets", planetListHandler).Queries("name", "{name}").Methods(http.MethodGet)
	r.HandleFunc("/api/planets/{id}", planetReadHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/planets/{id}", planetDeleteHandler).Methods(http.MethodDelete)

	http.Handle("/", r)

	if err := http.ListenAndServe("localhost:8081", nil); err != nil {
		fmt.Println(err)
	}
}