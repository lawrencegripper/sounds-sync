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
	var searchTerm, playlistID string
	flag.StringVar(&searchTerm, "show", "", "Enter a search term for the show you wish to sync")
	flag.StringVar(&playlistID, "playlist", "", "Enter the playlist ID from apple music to sync to")
	flag.Parse()

	// Search for the program
	searchUrl := fmt.Sprintf("https://rms.api.bbc.co.uk/v2/programmes/search/container?q=%s&experience=domestic", url.QueryEscape(searchTerm))
	searchResults := generated.SoundsProgramSearchResult{}
	err := fetchSoundsResponse(searchUrl, &searchResults)
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

	appleMusicTrackUrls := []string{}
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
				appleMusicTrackUrls = append(appleMusicTrackUrls, musicServiceUrl.Uri)
			}
		}
	}

	// Lets put these into the apple playlist
	addTracksRequest := ApplePlaylistTracksRequest{
		Tracks: AppleTracks{
			Data: []AppleTrack{},
		},
	}
	for _, trackUrl := range appleMusicTrackUrls {
		parsedUrl, err := url.Parse(trackUrl)
		if err != nil {
			panic(err)
		}
		parsedQuery, err := url.ParseQuery(parsedUrl.RawQuery)
		if err != nil {
			panic(err)
		}
		addTracksRequest.Tracks.Data = append(addTracksRequest.Tracks.Data, AppleTrack{
			ID: parsedQuery["i"][0],
		})
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

type ApplePlaylistTracksRequest struct {
	Tracks AppleTracks `json:"tracks"`
}

type AppleTracks struct {
	Data []AppleTrack `json:"data"`
}

type AppleTrack struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
