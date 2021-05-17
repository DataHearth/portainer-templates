BINARY_FOLDER = bin/portainer-templates

build:
	@command -v go >/dev/null || (echo 'go CLI is required to build portainer-templates'; exit 1)
	@echo "Create portainer-templates binary"
	go build -o $(BINARY_FOLDER) cmd/main.go
	@echo "Binary created in $(BINARY_FOLDER)"

create-docker-image:
	@command -v docker >/dev/null || (echo 'docker is required to create docker image'; exit 1)
	@read -p "Docker release: " release; \
    docker build --tag datahearth/portainer-templates:$$release .

push-docker-image:
	@command -v docker >/dev/null || (echo 'docker is required to create docker image'; exit 1)
	@read -p "Docker image tag: " tag; \
    docker push datahearth/portainer-templates:$$tag .
