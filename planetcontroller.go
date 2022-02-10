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

	if name == "" || climate == "" || terrain == "" {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	planet, err := createPlanet(name, climate, terrain)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(planet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", ContentTypeDefault)
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}

func planetSearchHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	planets, err := searchPlanets(vars["name"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(planets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", ContentTypeDefault)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func planetListHandler(w http.ResponseWriter, r *http.Request) {

	planets, err := listPlanets()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(planets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", ContentTypeDefault)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func planetReadHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	planet, err := getPlanetById(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(planet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", ContentTypeDefault)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func planetDeleteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	_, err := deletePlanetById(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", ContentTypeDefault)
	w.WriteHeader(http.StatusOK)
	//w.Write(data)
}
