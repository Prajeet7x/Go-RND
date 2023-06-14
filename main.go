package main

import goapi "go-api/API"

func main() {
	URL := "https://swapi.dev/api/people/"

	goapi.StarWarsAPI(URL)
}
