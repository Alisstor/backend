# Backend

A Go backend starter using Gin, GORM, MySQL, JWT auth, and environment-based configuration.

## Prerequisites

- Go 1.22+
- MySQL server running locally or available on your machine

## Setup

1. Copy the environment example file:

```bash
cp .env.example .env
```

2. Update the values in `.env` for your local database and JWT settings.

3. Install dependencies:

```bash
go mod tidy
```

4. Run the backend:

```bash
go run ./cmd/api
```

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

## Default routes

- `GET /health`
- `GET /ping`
- `GET /api/protected`

## Notes

- This project is intentionally minimal and ready for extension.
- Local development is expected to use your own database/server setup rather than Docker-based setup files.
