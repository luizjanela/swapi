package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Starting SWAPI Server... May the force be with us!")

	r := mux.NewRouter()

	r.HandleFunc("/api/planets", planetCreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/planets", planetSearchHandler).Queries("name", "{name}").Methods(http.MethodGet)
	r.HandleFunc("/api/planets", planetListHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/planets/{id}", planetReadHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/planets/{id}", planetDeleteHandler).Methods(http.MethodDelete)

	http.Handle("/", r)

	if err := http.ListenAndServe(ApiEndpoint, nil); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server C3P0 is running at http://%s\n", ApiEndpoint)
}
