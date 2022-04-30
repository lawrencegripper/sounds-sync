-include .env
DEV_CONTAINER_TAG:=ghcr.io/lawrencegripper/sounds-sync/devcontainer:latest

build:
	go build .

run-sync: build
	APPLE_BEARER_TOKEN="${APPLE_BEARER_TOKEN}" APPLE_USER_TOKEN="${APPLE_USER_TOKEN}" ./scripts/sync-playlists.sh

gen_models:
	./scripts/gen-models.sh

## devcontainer:
##		Builds the devcontainer used for VSCode and CI
devcontainer-build:
	@echo "Building devcontainer using tag: $(DEV_CONTAINER_TAG)"
	docker build -f ./.devcontainer/Dockerfile ./.devcontainer -t $(DEV_CONTAINER_TAG)

## devcontainer-run:
##		Runs the CMD variable in the devcontainer, eg: CMD='ruby ${PWD}/scripts/release.rb' make devcontainer-run 
devcontainer-run-sync: devcontainer-build
	@echo "Using tag: $(DEV_CONTAINER_TAG)"
	# Note command mirrors required envs from host into container. Using '@' to avoid showing values in CI output.
	@docker run -v ${PWD}:${PWD} \
		-e APPLE_BEARER_TOKEN="$(APPLE_BEARER_TOKEN)" \
		-e APPLE_USER_TOKEN="$(APPLE_USER_TOKEN)" \
		--entrypoint /bin/bash \
		--workdir "${PWD}" \
		$(DEV_CONTAINER_TAG) \
		-c "./scripts/sync-playlists.sh"
