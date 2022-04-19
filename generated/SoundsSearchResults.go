package generated

type SoundsSearchResults struct {
	Data []struct {
		Controls interface{} `json:"controls"`
		Data     []struct {
			Activities []struct {
				Action string `json:"action"`
				Type   string `json:"type"`
				Urn    string `json:"urn"`
			} `json:"activities"`
			Id        string `json:"id"`
			Image_url string `json:"image_url"`
			Network   struct {
				Id          string `json:"id"`
				Key         string `json:"key"`
				Logo_url    string `json:"logo_url"`
				Short_title string `json:"short_title"`
			} `json:"network"`
			Playable_count interface{} `json:"playable_count"`
			Synopses       struct {
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
		Description interface{} `json:"description"`
		Id          string      `json:"id"`
		Image_url   interface{} `json:"image_url"`
		State       string      `json:"state"`
		Style       interface{} `json:"style"`
		Title       string      `json:"title"`
		Total       interface{} `json:"total"`
		Type        string      `json:"type"`
		Uris        interface{} `json:"uris"`
	} `json:"data"`
}

