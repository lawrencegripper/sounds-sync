include .env

build:
	go build .

run: build
	@./sounds-sync -apple-playlist "${APPLE_PLAYLIST}" -show "${SHOW_NAME}" -apple-bearer-token "${APPLE_BEARER_TOKEN}" -apple-user-token "${APPLE_USER_TOKEN}"

gen_models:
	./scripts/gen-models.sh