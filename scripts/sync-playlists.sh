#!/bin/bash
set -e
echo "Building"
make build 

echo "Starting sync"
./sounds-sync -apple-playlist "p.8Wx6vrDfV2R7oZZ" -show "Radio 1's Chillest Show" -apple-bearer-token "${APPLE_BEARER_TOKEN}" -apple-user-token "${APPLE_USER_TOKEN}"
./sounds-sync -apple-playlist "p.zp6KqBpcmNJl200" -show "The Craig Charles Funk and Soul Show" -apple-bearer-token "${APPLE_BEARER_TOKEN}" -apple-user-token "${APPLE_USER_TOKEN}"
./sounds-sync -apple-playlist "p.oOzApM7flzmLxDD" -show "Radio 1's Indie Show" -apple-bearer-token "${APPLE_BEARER_TOKEN}" -apple-user-token "${APPLE_USER_TOKEN}"