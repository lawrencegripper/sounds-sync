# How will this work?

Swagger API Docs: https://rms.api.bbc.co.uk/docs/

## Find the show ID

https://rms.api.bbc.co.uk/v2/my/experience/inline/search?q=Chillest%20show

[response](./responses/search-results.json)

```
{
	"$schema": "https://rms.api.bbc.co.uk/docs/swagger.json#/definitions/ExperienceResponse",
	"data": [{
		"type": "inline_display_module",
		"id": "container_search",
		"style": null,
		"title": "Shows",
		"description": null,
		"state": "ok",
		"uris": null,
		"controls": null,
		"total": null,
		"data": [{
			"type": "container_item",
			"id": "b03hjfww",                             <------ Locate the ID of the show
			"urn": "urn:bbc:radio:brand:b03hjfww",
			"tlec_urn": "urn:bbc:radio:brand:b03hjfww",
			"network": {
				"id": "bbc_radio_one",
				"key": "radio1",
				"short_title": "Radio 1",
				"logo_url": "https://sounds.files.bbci.co.uk/2.3.0/networks/bbc_radio_one/{type}_{size}.{format}"
			},
			"titles": {
				"primary": "Radio 1's Chillest Show",
				"secondary": null,
				"tertiary": null
			},
```

## Find an instance of the show you want

https://rms.api.bbc.co.uk/v2/my/experience/inline/container/urn:bbc:radio:brand:b03hjfww?hero_enabled=true&tag_enabled=true

[response](./responses/instance-result.json)

```
{
                    "type": "playable_item",                  <------ Look for playable items
                    "id": "m0015s68",                         <------ Locate the ID for an instance of the show
                    "urn": "urn:bbc:radio:episode:m0015s69",
                    "network": {
                        "id": "bbc_radio_one",
                        "key": "radio1",
                        "short_title": "Radio 1",
                        "logo_url": "https://sounds.files.bbci.co.uk/2.3.0/networks/bbc_radio_one/{type}_{size}.{format}"
                    },
                    "titles": {
                        "primary": "Radio 1's Chillest Show",
                        "secondary": "Lola Young Chill Mix",
                        "tertiary": null
                    },
```

## Request the track listings for the show

https://rms.api.bbc.co.uk/v2/versions/m0015m41/segments?

[response](./responses/tracks.json)

Get the apple music ID from the uris

```
 {
            "type": "segment_item",
            "id": "p0bxvhhx",
            "urn": "urn:bbc:radio:segment:music:nznqj9",
            "segment_type": "music",
            "titles": {
                "primary": "Billie Holiday",
                "secondary": "I'll Be Seeing You",
                "tertiary": null
            },
            "synopses": null,
            "image_url": null,
            "offset": {
                "start": 4225,
                "end": 4431,
                "label": null,
                "now_playing": false
            },
            "uris": [
                {
                    "type": "commercial-music-service",
                    "id": "commercial-music-service-spotify",
                    "label": "Spotify",
                    "uri": "https://open.spotify.com/track/4smkJW6uzoHxGReZqqwHS5"
                },
                {
                    "type": "commercial-music-service",
                    "id": "commercial-music-service-apple",
                    "label": "Apple Music",
                    "uri": "https://music.apple.com/gb/album/ill-be-seeing-you/1443140316?i=1443141555"
                }
            ]
        },
```

## Get existing or Create a playlist in Apple Music

Get: https://developer.apple.com/documentation/applemusicapi/get_a_library_playlist

`GET https://api.music.apple.com/v1/me/library/playlists/{id}`

Create: https://developer.apple.com/documentation/applemusicapi/create_a_new_library_playlist?utm_source=pocket_mylist

`POST https://api.music.apple.com/v1/me/library/playlists`

## Add the Tracks

https://developer.apple.com/documentation/applemusicapi/add_tracks_to_a_library_playlist?utm_source=pocket_mylist

`POST https://api.music.apple.com/v1/me/library/playlists/{id}/tracks`