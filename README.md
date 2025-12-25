# Greenlight — Movie API

A small RESTful API written in Go to manage movies and users. Built with httprouter and a simple data layer. This README describes setup, running, available endpoints, and examples.

## Features

- CRUD for movies (create, list, show, update, delete)
- User registration endpoint
- JSON request/response handling and validation
- Clean project structure with `cmd/api`, `internal/data`, and helper packages

## Tech stack

- Go (1.20+)
- PostgreSQL
- httprouter
- Standard library + small helper packages (validator, custom JSON helpers)

## Project layout (high level)

- cmd/api — API server entrypoint and routes
- internal/data — models and DB interaction
- internal/validator — validation helpers
- migrations/ — SQL migration files (if present)
- README.md — this file

## Prerequisites

- Go (install from https://go.dev)
- PostgreSQL
- psql CLI (optional but useful)

## Environment / configuration

The server reads a DSN and a port from environment variables (adjust names if your code uses others). Example DSN format:

export DB_DSN="postgres://user:password@localhost/greenlight?sslmode=disable"
export PORT=4000

Set any other env vars your code expects (DB pool sizes, timeouts, etc.).

## Database setup (example)

Replace user/password/database names to match your environment.

1. Create database and user (example):

   1. Start psql as a superuser:
      psql -U postgres
   2. Create DB and user:
      CREATE DATABASE greenlight;
      CREATE USER greenlight_user WITH PASSWORD 'yourpassword';
      GRANT ALL PRIVILEGES ON DATABASE greenlight TO greenlight_user;
   3. Exit psql.

2. Run migrations:
   - If you have SQL migration files in a `migrations/` folder, apply them with your preferred tool (psql, goose, migrate, or a included migrations binary).
   - Example (simple, applying SQL files manually):
     psql "$DB_DSN" -f ./migrations/0001_init.sql

Make sure your `DB_DSN` points to the created DB and user.

## Build & run

From project root:

1. Download deps:
   go mod download

2. Run directly:
   PORT=4000 DB_DSN="postgres://user:pass@localhost/greenlight?sslmode=disable" go run ./cmd/api

3. Or build binary:
   go build -o bin/api ./cmd/api
   PORT=4000 DB_DSN="postgres://user:pass@localhost/greenlight?sslmode=disable" ./bin/api

Server listens on the port you set (default seen in examples: 4000).

## API Endpoints

Base path: /v1

- Healthcheck

  - GET /v1/healthcheck
  - Example:
    curl -i http://localhost:4000/v1/healthcheck

- Movies

  - GET /v1/movies

    - List movies (supports filters via query string — check code for supported params).
    - Example:
      curl -i "http://localhost:4000/v1/movies"

  - POST /v1/movies

    - Create a movie (JSON body).
    - Example:
      curl -i -X POST -H "Content-Type: application/json" \
       -d '{"title":"Moana","year":2016,"runtime":"107 mins","genres":["animation","adventure"]}' \
       http://localhost:4000/v1/movies

  - GET /v1/movies/:id

    - Retrieve single movie by ID.
    - Example:
      curl -i http://localhost:4000/v1/movies/1

  - PATCH /v1/movies/:id

    - Partial update. Send only the fields to update. Must send `Content-Type: application/json`.
    - Example (update year):
      curl -i -X PATCH -H "Content-Type: application/json" \
       -d '{"year":1985}' \
       http://localhost:4000/v1/movies/1
    - Note: do not call the JSON body twice — ensure your handler only decodes the request once.

  - DELETE /v1/movies/:id
    - Delete movie by ID.
    - Example:
      curl -i -X DELETE http://localhost:4000/v1/movies/1

- Users
  - POST /v1/users
    - Register a new user (see your user struct for required fields).
    - Example:
      curl -i -X POST -H "Content-Type: application/json" \
       -d '{"name":"Alice","email":"alice@example.com","password":"secret"}' \
       http://localhost:4000/v1/users

## JSON shape (movies)

Request body for creating a movie:
{
"title": "string",
"year": 2020,
"runtime": "120 mins",
"genres": ["genre1","genre2"]
}

For PATCH, use pointer fields (handled by your handler) — only include fields you want to update.

## Common troubleshooting

- "body must not be empty" when PATCHing:

  - Make sure to send `Content-Type: application/json` with curl (`-H "Content-Type: application/json"`).
  - Ensure your handler only calls the JSON decoding helper once — decoding twice will fail because the body stream is consumed.

- DB connection errors:
  - Confirm `DB_DSN` is correct and DB is reachable.
  - Ensure migrations were applied.

## Testing

Run unit tests:
go test ./...

If you have test DB setup required, configure environment variables for tests accordingly.

## Linting & formatting

- gofmt and go vet:
  go vet ./...
  gofmt -w .

---
