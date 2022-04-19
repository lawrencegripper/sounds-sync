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
	flag.StringVar(&searchTerm, "search term", "", "Enter a search term for the show you wish to sync")
	url := fmt.Sprintf("https://rms.api.bbc.co.uk/v2/my/experience/inline/search?q=%s", url.QueryEscape(searchTerm))
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	rawResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	searchResults := generated.SoundsSearchResults{}
	err = json.Unmarshal(rawResponse, &searchResults)

	if err != nil {
		panic(err)
	}
}
