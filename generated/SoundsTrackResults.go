package generated

type SoundsTrackResults struct {
	Data []struct {
		Id        string      `json:"id"`
		Image_url interface{} `json:"image_url"`
		Offset    struct {
			End         int         `json:"end"`
			Label       interface{} `json:"label"`
			Now_playing bool        `json:"now_playing"`
			Start       int         `json:"start"`
		} `json:"offset"`
		Segment_type string      `json:"segment_type"`
		Synopses     interface{} `json:"synopses"`
		Titles       struct {
			Primary   string      `json:"primary"`
			Secondary string      `json:"secondary"`
			Tertiary  interface{} `json:"tertiary"`
		} `json:"titles"`
		Type string `json:"type"`
		Uris []struct {
			Id    string `json:"id"`
			Label string `json:"label"`
			Type  string `json:"type"`
			Uri   string `json:"uri"`
		} `json:"uris"`
		Urn string `json:"urn"`
	} `json:"data"`
}

