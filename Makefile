# Go variables
GOCMD = go
GOBUILD = $(GOCMD) build
BINARY_NAME = go_api
DOCKER_IMAGE_NAME = my-go-api
DOCKER_CONTAINER_NAME = my-go-api-container
SERVER_PORT = 8080

# Default target
.DEFAULT_GOAL := run

# Targets
build:
	$(GOBUILD) -o bin/$(BINARY_NAME) ./cmd/api

deps:
	$(GOCMD) mod tidy

run: build
	./bin/$(BINARY_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run -d -p $(SERVER_PORT):$(SERVER_PORT) --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME)

docker-stop:
	docker stop $(DOCKER_CONTAINER_NAME)

docker-rm:
	docker rm $(DOCKER_CONTAINER_NAME)

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build         : Build the application"
	@echo "  deps          : Install dependencies"
	@echo "  run           : Build and run the application locally"
	@echo "  docker-build  : Build Docker image"
	@echo "  docker-run    : Run Docker container detached"
	@echo "  docker-stop   : Stop Docker container"
	@echo "  docker-rm     : Remove Docker container"
	@echo ""
	@echo "To start the application locally:"
	@echo "  make run"
	@echo ""
	@echo "To start the Docker container:"
	@echo "  make docker-run"
	@echo ""
	@echo "To stop the Docker container:"
	@echo "  make docker-stop"