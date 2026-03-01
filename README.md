# Microservices & Front-End Deployment

This repo contains a Go-based microservices stack and a front-end app. Use the **Makefile** in `project/` to build and run everything.

## Prerequisites

- **Docker & Docker Compose** – for running services in containers
- **Go 1.21+** – for building binaries locally (optional)
- **Make** – for running the commands below

## Quick Start (Docker Compose)

From the **project** directory, a single command will build and run **both** the back-end services and the front-end behind Caddy:

```bash
cd project
make up_build
```

On **Windows**, use the Windows-specific Makefile:

```bash
cd project
make -f Makefile.windows up_build
```

This will:

1. Build Linux binaries for broker, auth, logger, mail, listener, and front-end.
2. Bring down any existing stack
3. Build and start all Docker Compose services (front-end, Caddy, broker, auth, logger, mailer, listener, Postgres, Mongo, RabbitMQ, MailHog)

Once the stack is running, you can access the application at `http://localhost:80`.

To start existing images without rebuilding (e.g. after a reboot), you can use:

```bash
cd project
make up
```

To stop the stack:

```bash
cd project
make down
```

## Makefile Targets (run from `project/`)

| Command        | Description |
|----------------|-------------|
| `make up`      | Start all containers in the background (no rebuild) |
| `make up_build` | Build all service binaries, then build and start all containers |
| `make down`    | Stop and remove Docker Compose containers |
| `make build_broker`   | Build broker-service Linux binary only |
| `make build_auth`     | Build authentication-service Linux binary only |
| `make build_logger`   | Build logger-service Linux binary only |
| `make build_mail`     | Build mail-service Linux binary only |
| `make build_listener` | Build listener-service Linux binary only |
| `make build_front`    | Build front-end binary (current OS) |
| `make start`   | Build and run the front-end app (e.g. `./frontApp`) in the background |
| `make stop`    | Stop the front-end app process |

## Service Ports (when using Docker Compose)

- **Broker**: `8080` → 80 (internal)
- **Authentication**: `8081` → 80 (internal)
- **Postgres**: `5432`
- **Mongo**: `27017`
- **RabbitMQ**: `5672`
- **MailHog SMTP**: `1025` (API/ui: `8025`)

## Front-End (standalone)

To build and run the front-end **without Docker** (pure Go binary on your host):

```bash
cd project
make build_front
make start
```

To stop it:

```bash
make stop
```

The front-end serves on port 80 by default.

> **Note:** When using Docker Compose, you do **not** need to run any additional `make` command to start the front-end separately – it is started automatically as part of `make up_build` / `make up`.

## Docker Swarm (optional)

You can also run the stack using **Docker Swarm** with the `swarm.yml` file in the `project/` directory.

From the **project** directory:

```bash
cd project

# 1. Initialize Swarm (only once per Docker host)
docker swarm init

# 2. Deploy the stack using swarm.yml
docker stack deploy -c swarm.yml microservices

# 3. Scale a service (example: 3 replicas of broker-service)
docker service scale microservices_broker-service=3

# 4. Update a service to use a new image tag
docker service update --image your-registry/your-image:latest microservices_broker-service

# 5. Remove the deployed stack
docker stack rm microservices
```

When the Swarm stack is running, the application is available at `http://localhost:80`.

## Project Layout

- **project/** – Docker Compose, Makefile, and db-data volumes
- **broker-service/** – API gateway / broker
- **authentication-service/** – User auth (Postgres)
- **logger-service/** – Log storage (MongoDB, RPC, gRPC)
- **mail-service/** – SMTP mail sender
- **listener-service/** – RabbitMQ consumer (log/auth events)
- **front-end/** – Web UI

Each service has its own README with a description and code structure.
