# Makefile for Docker Compose setup

# Default target
.PHONY: help
help:
	@echo "Makefile for managing Docker containers"
	@echo ""
	@echo "Available targets:"
	@echo "  build       - Build the Docker images"
	@echo "  up          - Start the containers in the background"
	@echo "  up-d         - Start the containers in detached mode"
	@echo "  down        - Stop and remove the containers"
	@echo "  restart     - Restart the containers"
	@echo "  logs        - View the logs of all containers"
	@echo "  rebuild     - Rebuild the Docker images"
	@echo "  rebuild-app - Rebuild only the app container"
	@echo "  clean       - Stop and remove the containers, networks, and volumes"
	@echo "  ps          - List running containers"

# Build the Docker images
.PHONY: build
build:
	docker-compose build

# Start the containers in the background
.PHONY: up
up:
	docker-compose up

# Start the containers in detached mode
.PHONY: up-d
up-d:
	docker-compose up -d

# Stop and remove the containers
.PHONY: down
down:
	docker-compose down

# Restart the containers
.PHONY: restart
restart: down up

# View the logs of all containers
.PHONY: logs
logs:
	docker-compose logs -f

# Rebuild the Docker images
.PHONY: rebuild
rebuild:
	docker-compose down --volumes --rmi all
	docker-compose build
	docker-compose up -d

# Rebuild only the app container
.PHONY: rebuild-app
rebuild-app:
	docker-compose build app
	docker-compose up -d app

# Stop and remove the containers, networks, and volumes
.PHONY: clean
clean:
	docker-compose down --volumes --remove-orphans

# List running containers
.PHONY: ps
ps:
	docker-compose ps
