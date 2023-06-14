package goapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type star_wars struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:results`
}

//This function demonstrates how to handle public APIs and unmarshal the data
func StarWarsAPI(PUBLIC_API_URL string) {

	req, err := http.NewRequest("GET", PUBLIC_API_URL, nil)
	if err != nil {
		fmt.Println("Error catching request: ", err)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading the response body:", res)
		return
	}

	// Unmarshal the JSON response object into a struct
	var starWarsResponse star_wars
	err = json.Unmarshal(body, &starWarsResponse)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON:", err)
		return
	}

	names := make([]string, 0)
	for _, character := range starWarsResponse.Results {
		names = append(names, character.Name)
	}

	fmt.Println("Names of the star wars characters:")
	for _, name := range names {
		fmt.Println(name)
	}
}
