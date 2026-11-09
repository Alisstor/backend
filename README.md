# Go Backend Starter

A minimal backend setup using Go, Gin, GORM, PostgreSQL, JWT, and environment-based configuration.

## Prerequisites

- Go 1.22+
- Docker and Docker Compose
- PostgreSQL (or use the included Docker service)

## Setup

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

2. Start PostgreSQL with Docker:
   ```bash
   docker compose up -d postgres
   ```

3. Install module dependencies:
   ```bash
   go mod tidy
   ```

4. Run the app:
   ```bash
   go run ./cmd/api
   ```

## Default endpoints

- `GET /health` - health check
- `GET /ping` - simple ping endpoint
- `GET /api/protected` - authenticated route using JWT bearer token

## Environment variables

- `SERVER_PORT`
- `DB_HOST`
- `DB_PORT`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `DB_SSLMODE`
- `JWT_SECRET`
- `JWT_TTL_MINUTES`

## Notes

- The project is ready to extend with controllers, services, repositories, and real JWT-auth user flows.
- `.env` is ignored by Git to keep local secrets out of source control.
