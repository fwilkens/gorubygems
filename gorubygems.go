package gorubygems

import (
	"encoding/json"
	"net/http"
)

type Gem struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	SourceCode string `json:"source_code_uri"`
	Homepage   string `json:"homepage_uri"`
	Platform   string `json:"platform"`
}

const baseUrl = "https://rubygems.org/api/v1"
const RequestType = ".json"

func JustUpdated() []Gem {
	return get("/activity/just_updated", "")
}

func Search(term string) []Gem {
	return get("/search", "?query="+term)
}

func Latest() []Gem {
	return get("/activity/latest", "")
}

func get(endpoint string, queryString string) []Gem {
	resourceUri := baseUrl + endpoint + RequestType + queryString
	response, err := http.Get(resourceUri)
	checkError(err)

	defer response.Body.Close()
	var gems []Gem
	json.NewDecoder(response.Body).Decode(&gems)
	return gems
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
