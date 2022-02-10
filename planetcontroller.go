package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func planetCreateHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	climate := r.FormValue("climate")
	terrain := r.FormValue("terrain")

	// Todo validate each of them separately
	// Validate climate and terrain options
	if name == "" || climate == "" || terrain == "" {
		w.WriteHeader(http.StatusPreconditionFailed)
		data, _ := json.Marshal(ResponseModel{false, ResponseErrorAllParamsMustBeSet})
		w.Write(data)
		return
	}

	// Get apparitions using SW Library
	apparitions, err := getApparitionsByName(name)

	// Check if planet name already exists using ELK Library
	planet, err := searchPlanetByName(name)

	// A planet with the same name already exists
	// Todo implementar regra de lowercase para casos de diferença apenas de case sensitive
	if planet.Name != "" {
		w.WriteHeader(http.StatusConflict)
		data, _ := json.Marshal(ResponseModel{false, ResponseErrorPlanetAlreadyExists})
		w.Write(data)
		return
	}

	if err.Error() != ResponseErrorNoPlanetFound {
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := json.Marshal(ResponseModel{false, err.Error()})
		w.Write(data)
		return
	}

	// ELK Library
	// Planet created with empty id
	planet, err = createPlanet(PlanetModel{"", name, climate, terrain, apparitions})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := json.Marshal(ResponseModel{false, err.Error()})
		w.Write(data)
		return
	}

	data, err := json.Marshal(planet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := json.Marshal(ResponseModel{false, err.Error()})
		w.Write(data)
		return
	}

	w.Header().Add("Content-Type", ContentTypeDefault)
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}

func planetSearchHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// ELK Library
	planet, err := searchPlanetByName(vars["name"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		data, _ := json.Marshal(ResponseModel{false, ResponseErrorPlanetNotFound})
		w.Write(data)
		return
	}

	data, err := json.Marshal(planet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := json.Marshal(ResponseModel{false, err.Error()})
		w.Write(data)
		return
	}

	w.Header().Add("Content-Type", ContentTypeDefault)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func planetListHandler(w http.ResponseWriter, r *http.Request) {

	// Todo implement pagination logic (limit, offset/start)
	// ELK Library
	planets, err := listPlanets()

	if len(planets) == 0 {
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal([]string{})
		w.Write(data)
		return
	}

	data, err := json.Marshal(planets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := json.Marshal(ResponseModel{false, err.Error()})
		w.Write(data)
		return
	}

	w.Header().Add("Content-Type", ContentTypeDefault)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func planetReadHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// ELK Library
	planet, err := getPlanetById(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		data, _ := json.Marshal(ResponseModel{false, ResponseErrorPlanetNotFound})
		w.Write(data)
		return
	}

	data, err := json.Marshal(planet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := json.Marshal(ResponseModel{false, err.Error()})
		w.Write(data)
		return
	}

	w.Header().Add("Content-Type", ContentTypeDefault)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func planetDeleteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// ELK Library
	_, err := deletePlanetById(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		data, _ := json.Marshal(ResponseModel{false, ResponseErrorPlanetNotFound})
		w.Write(data)
		return
	}

	w.Header().Add("Content-Type", ContentTypeDefault)
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(ResponseModel{true, ResponsePlanetDeleted})
	w.Write(data)
}
