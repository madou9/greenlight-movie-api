# Greenlight — Movie API in Go

Greenlight is a JSON API for retrieving and managing movie information, built in Go with a focus on clean architecture, performance, and production-ready practices.

The API exposes endpoints for movies, users, authentication tokens, and application health/metrics. PostgreSQL is used for persistent storage, and the service is intended to be deployed on a Linux server environment.

---

## 🚀 Planned Features (End Goals)

| Method | Route                     | Description                       |
| ------ | ------------------------- | --------------------------------- |
| GET    | /v1/healthcheck           | Show application status & version |
| GET    | /v1/movies                | List all movies                   |
| POST   | /v1/movies                | Create a new movie                |
| GET    | /v1/movies/:id            | Get a specific movie              |
| PATCH  | /v1/movies/:id            | Update a specific movie           |
| DELETE | /v1/movies/:id            | Delete a movie                    |
| POST   | /v1/users                 | Register a user                   |
| PUT    | /v1/users/activated       | Activate a user                   |
| PUT    | /v1/users/password        | Change a user's password          |
| POST   | /v1/tokens/authentication | Generate auth token               |
| POST   | /v1/tokens/password-reset | Generate password reset token     |
| GET    | /debug/vars               | Show metrics                      |

The API will use **PostgreSQL** for persistent storage and will eventually be deployed on a Linux server (DigitalOcean).

---

## 📚 Purpose of This Repository

This repo is not just the final code — it documents my learning journey:

- Every step will be committed with meaningful messages.
- I will write short notes or comments as I learn new techniques.
- The goal is to demonstrate **growth and understanding**, not just the end result.

---

## 🛠️ Tech Stack

- **Language:** Go
- **Database:** PostgreSQL
- **Auth:** Token-based authentication (JWT-like)
- **Deployment target:** Linux VM on DigitalOcean

---

## ⚙️ Running the Project Locally

### 2. Start PostgreSQL Locally

You must have **PostgreSQL installed** on your machine.

#### **Create the database:**

```bash
createdb greenlight

## Run the migrations:
migrate -path ./migrations -database "postgres://localhost/greenlight?sslmode=disable" up

## Export ENV variables
export GREENLIGHT_DB_DSN="postgres://YOURUSER@localhost/greenlight?sslmode=disable"
export GREENLIGHT_PORT=4000

### run the API
    - go run ./cmd/api/

### TEST Request

 BODY='{"title":"Deadpool","year":2016,"runtime":"108 mins","genres":["action","comedy"]}'

curl -d "$BODY" \
     -H "Content-Type: application/json" \
     localhost:4000/v1/movies

```

## 👤 Author

Learning and built by **Hama Issoufou**

---

```

```
