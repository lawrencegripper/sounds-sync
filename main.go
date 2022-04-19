package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	var searchTerm string
	flag.StringVar(&searchTerm, "search term", "", "Enter a search term for the show you wish to sync")
	url := fmt.Sprintf("https://rms.api.bbc.co.uk/v2/my/experience/inline/search?q=%s", url.QueryEscape(searchTerm))
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
}
