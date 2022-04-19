package generated

type SoundsProgramSearchResult struct {
	Data []struct {
		Activities []interface{} `json:"activities"`
		Id         string        `json:"id"`
		Image_url  string        `json:"image_url"`
		Network    struct {
			Id          string `json:"id"`
			Key         string `json:"key"`
			Logo_url    string `json:"logo_url"`
			Short_title string `json:"short_title"`
		} `json:"network"`
		Synopses struct {
			Long   interface{} `json:"long"`
			Medium string      `json:"medium"`
			Short  string      `json:"short"`
		} `json:"synopses"`
		Titles struct {
			Primary   string      `json:"primary"`
			Secondary interface{} `json:"secondary"`
			Tertiary  interface{} `json:"tertiary"`
		} `json:"titles"`
		Tlec_urn string `json:"tlec_urn"`
		Type     string `json:"type"`
		Uris     []struct {
			Id    interface{} `json:"id"`
			Label string      `json:"label"`
			Type  string      `json:"type"`
			Uri   string      `json:"uri"`
		} `json:"uris"`
		Urn string `json:"urn"`
	} `json:"data"`
}

