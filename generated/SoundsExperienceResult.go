package generated

import "encoding/json"

type SoundsExperienceResult struct {
	Data []struct {
		Controls    interface{}     `json:"controls"`
		Data        json.RawMessage `json:"data"`
		Description interface{}     `json:"description"`
		Id          string          `json:"id"`
		Image_url   interface{}     `json:"image_url"`
		State       string          `json:"state"`
		Style       interface{}     `json:"style"`
		Title       string          `json:"title"`
		Total       int             `json:"total"`
		Type        string          `json:"type"`
		Uris        interface{}     `json:"uris"`
	} `json:"data"`
}

type SoundExperienceEpisodeData struct {
	Activities   []interface{} `json:"activities"`
	Availability struct {
		From  string `json:"from"`
		Label string `json:"label"`
		To    string `json:"to"`
	} `json:"availability"`
	Container struct {
		Activities []interface{} `json:"activities"`
		Id         string        `json:"id"`
		Synopses   struct {
			Long   interface{} `json:"long"`
			Medium string      `json:"medium"`
			Short  string      `json:"short"`
		} `json:"synopses"`
		Title string `json:"title"`
		Type  string `json:"type"`
		Urn   string `json:"urn"`
	} `json:"container"`
	Download struct {
		Quality_variants struct {
			High struct {
				Bitrate   int    `json:"bitrate"`
				File_size int    `json:"file_size"`
				File_url  string `json:"file_url"`
				Label     string `json:"label"`
			} `json:"high"`
			Low struct {
				Bitrate   int    `json:"bitrate"`
				File_size int    `json:"file_size"`
				File_url  string `json:"file_url"`
				Label     string `json:"label"`
			} `json:"low"`
			Medium struct {
				Bitrate   int    `json:"bitrate"`
				File_size int    `json:"file_size"`
				File_url  string `json:"file_url"`
				Label     string `json:"label"`
			} `json:"medium"`
		} `json:"quality_variants"`
		Type string `json:"type"`
	} `json:"download"`
	Duration struct {
		Label string `json:"label"`
		Value int    `json:"value"`
	} `json:"duration"`
	Guidance struct {
		Competition_warning bool        `json:"competition_warning"`
		Warnings            interface{} `json:"warnings"`
	} `json:"guidance"`
	Id        string `json:"id"`
	Image_url string `json:"image_url"`
	Network   struct {
		Id          string `json:"id"`
		Key         string `json:"key"`
		Logo_url    string `json:"logo_url"`
		Short_title string `json:"short_title"`
	} `json:"network"`
	Play_context   interface{} `json:"play_context"`
	Progress       interface{} `json:"progress"`
	Recommendation interface{} `json:"recommendation"`
	Release        struct {
		Date  string `json:"date"`
		Label string `json:"label"`
	} `json:"release"`
	Synopses struct {
		Long   string `json:"long"`
		Medium string `json:"medium"`
		Short  string `json:"short"`
	} `json:"synopses"`
	Titles struct {
		Primary   string      `json:"primary"`
		Secondary string      `json:"secondary"`
		Tertiary  interface{} `json:"tertiary"`
	} `json:"titles"`
	Type string `json:"type"`
	Uris []struct {
		Id    interface{} `json:"id"`
		Label string      `json:"label"`
		Type  string      `json:"type"`
		Uri   string      `json:"uri"`
	} `json:"uris"`
	Urn string `json:"urn"`
}
