# Sync BBC Sounds track listing to Apple Music

Given a BBC Sounds Radio show this utility finds the tracks and syncs them to an Apple Music Playlist

Search "SoundsSync" on Apple Music to see the playlists, they're updated at midnight UTC. 

Raise an issue to request a new show is added.

## Dev

Create a `.env` file with: 

```
APPLE_BEARER_TOKEN=
APPLE_USER_TOKEN=
APPLE_PLAYLIST=
SHOW_NAME=radio 1 chillest show
```

Get these by looking at the devtools of your browser while logged into `music.apple.com`.

Run `make run` and watch the songs get synced.