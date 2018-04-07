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
	response := get("/activity/just_updated", "")
	return decodeArray(response)
}

func Search(term string) []Gem {
	response := get("/search", "?query="+term)
	return decodeArray(response)
}

func Latest() []Gem {
	response := get("/activity/latest", "")
	return decodeArray(response)
}

func GemInfo(gemName string) Gem {
	response := get("/gems/"+gemName, "")
	return decodeObject(response)
}

func get(endpoint string, queryString string) *http.Response {
	resourceUri := baseUrl + endpoint + RequestType + queryString
	response, err := http.Get(resourceUri)
	checkError(err)
	return response
}

func decodeArray(response *http.Response) []Gem {
	defer response.Body.Close()
	var gems []Gem
	json.NewDecoder(response.Body).Decode(&gems)
	return gems
}

func decodeObject(response *http.Response) Gem {
	defer response.Body.Close()
	var gem Gem
	json.NewDecoder(response.Body).Decode(&gem)
	return gem
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
