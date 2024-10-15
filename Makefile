# Variables
DOCKER_COMPOSE_FILE = docker-compose.yml
APP_CONTAINER = app

# Default: Start Docker Compose
.PHONY: up
up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop Docker Compose
.PHONY: down
down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Restart the app container
.PHONY: restart-app
restart-app:
	docker-compose -f $(DOCKER_COMPOSE_FILE) restart $(APP_CONTAINER)

# View logs for the app container
.PHONY: logs
logs:
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f $(APP_CONTAINER)

# Tail logs for all containers
.PHONY: logs-all
logs-all:
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

# Build/rebuild containers (if needed)
.PHONY: build
build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

# Check container statuses
.PHONY: ps
ps:
	docker-compose -f $(DOCKER_COMPOSE_FILE) ps

# Clean up and remove all containers, networks, and volumes
.PHONY: clean
clean:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down -v
