package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/lawrencegripper/sounds-sync/generated"
)

func main() {
	var searchTerm string
	flag.StringVar(&searchTerm, "show", "", "Enter a search term for the show you wish to sync")
	flag.Parse()

	url := fmt.Sprintf("https://rms.api.bbc.co.uk/v2/programmes/search/container?q=%s&experience=domestic", url.QueryEscape(searchTerm))
	searchResults := generated.SoundsProgramSearchResult{}

	fmt.Println(url)
	err := fetchSoundsResponse(url, &searchResults)
	if err != nil {
		panic(err)
	}

	programId := searchResults.Data[0].Id
	fmt.Println(programId)
}

// This doesn't need generics, just rying it out
func fetchSoundsResponse(url string, response *generated.SoundsProgramSearchResult) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	rawResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(rawResponse, response)
}
