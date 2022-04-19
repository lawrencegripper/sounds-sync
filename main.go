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

	// Search for the program
	url := fmt.Sprintf("https://rms.api.bbc.co.uk/v2/programmes/search/container?q=%s&experience=domestic", url.QueryEscape(searchTerm))
	searchResults := generated.SoundsProgramSearchResult{}
	err := fetchSoundsResponse(url, &searchResults)
	if err != nil {
		panic(err)
	}

	// Take the top result
	programURN := searchResults.Data[0].Urn
	fmt.Println(programURN)

	// Look for episodes
	programInstanceUrl := fmt.Sprintf("https://rms.api.bbc.co.uk/v2/experience/inline/container/%s", programURN)
	programInstanceResults := generated.SoundsExperienceResult{}
	err = fetchSoundsResponse(programInstanceUrl, &programInstanceResults)
	if err != nil {
		panic(err)
	}

	tracks := []string{}
	// Note: The API json can vary types between struct and arrays for the same field name
	//       to workaround that we use the `.json.RawResponse` type and index ignore the asymetric item at index 0
	// First item is usually the display header this has data field as struct
	// todo: filter on `Type` field for `inline_display_module`
	// todo: nested loops ewww, refactor with gofunc or methods
	episodesDataField := programInstanceResults.Data[1].Data
	episodesData := []generated.SoundExperienceEpisodeData{}
	json.Unmarshal(episodesDataField, &episodesData)
	for _, episode := range episodesData {
		if episode.Type != "playable_item" {
			continue
		}

		fmt.Printf("Looking at episode: %s\n", episode.Titles.Primary)

		showUrl := fmt.Sprintf("https://rms.api.bbc.co.uk/v2/versions/%s/segments?", episode.Id)
		trackResults := generated.SoundsTrackResults{}
		err = fetchSoundsResponse(showUrl, &trackResults)
		if err != nil {
			panic(err)
		}

		for _, track := range trackResults.Data {
			for _, musicServiceUrl := range track.Uris {
				if musicServiceUrl.Label != "Apple Music" {
					continue
				}

				fmt.Printf("-- Adding track: %s - %s\n", track.Titles.Primary, track.Titles.Secondary)
				tracks = append(tracks, musicServiceUrl.Uri)
			}
		}
	}
}

func fetchSoundsResponse(url string, response interface{}) error {
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
