# secret-santa-backend

# Docker Compose Project with Makefile

This project simplifies the management of your Docker Compose setup using a `Makefile`. You can easily start, stop, restart containers, view logs, and perform other common tasks using simple `make` commands.

## Prerequisites

- **Docker**: Ensure that Docker is installed and running on your system. [Install Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: Ensure Docker Compose is installed. [Install Docker Compose](https://docs.docker.com/compose/install/)
- **Make**: Make sure you have `make` installed on your system. It typically comes pre-installed on Unix-based systems like Linux and macOS.

## Usage

The following `make` commands are available for managing your Docker Compose project:

### 1. Start Docker Compose

To start the services defined in your `docker-compose.yml` file:

```bash
make up
```

This will start the containers in detached mode (`-d`), so they run in the background.

### 2. Stop and Remove Docker Compose Containers

To stop and remove all containers, networks, and volumes:

```bash
make down
```

### 3. Restart the Application Container

To restart just the `app` container (assuming your Go app container is named `app`):

```bash
make restart-app
```

This command is useful when you want to quickly restart the application without affecting other services.

### 4. View Application Logs

To view the logs for the `app` container:

```bash
make logs
```

This will display the logs of the `app` container and keep streaming new log entries.

### 5. View Logs for All Containers

To view logs for all running containers:

```bash
make logs-all
```

This shows and streams logs from all services defined in your `docker-compose.yml`.

### 6. Build/Rebuild Containers

To force a rebuild of the containers (useful when you change Dockerfiles or dependencies):

```bash
make build
```

### 7. Check Container Statuses

To see the current status of all containers:

```bash
make ps
```

This will list the containers with their current status (running, exited, etc.).

### 8. Clean Up (Remove All Containers, Networks, and Volumes)

To stop and remove all containers, networks, and volumes created by Docker Compose:

```bash
make clean
```

### Example Workflow

1. **Start services**: `make up`
2. **Check logs**: `make logs`
3. **Restart the app container**: `make restart-app`
4. **Stop all services**: `make down`

## Customization

- The `Makefile` assumes the `app` container is the one you want to restart and check logs for. If your container is named differently, modify the `APP_CONTAINER` variable at the top