package main

type PlanetModel struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Climate         string `json:"climate"`
	Terrain         string `json:"terrain"`
	FilmApparitions string `json:"film_apparitions"`
}
