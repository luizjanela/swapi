package main

const ApiEndpoint = "localhost:8083"

const ElkHost = "http://127.0.0.1:9200/index-swapi-0001"
const SwHost = "https://swapi.dev"

const ContentTypeDefault = "application/json; charset=utf-8"

// Responses

const ResponseErrorAllParamsMustBeSet = "All params must be set"
const ResponseErrorPlanetAlreadyExists = "Planet already exists"
const ResponseErrorNoPlanetsFound = "No planets found"
const ResponseErrorPlanetNotFound = "Planet not found"
const ResponsePlanetDeleted = "Planet deleted"
