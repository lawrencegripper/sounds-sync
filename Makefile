build:
	go build .

run: build
	./sounds-sync -playlist "p.8Wx6vrDfV2R7oZZ" -show "chillest show"

gen_models:
	./scripts/gen-models.sh