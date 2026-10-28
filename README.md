# 🚀 Golang REST API

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue.svg?logo=go)](https://go.dev/)
[![Framework](https://img.shields.io/badge/Framework-Gin-00ADD8.svg?logo=gin)](https://github.com/gin-gonic/gin)
[![ORM](https://img.shields.io/badge/ORM-GORM-7952B3.svg)](https://gorm.io/)
[![Database](https://img.shields.io/badge/Database-PostgreSQL-4169E1.svg?logo=postgresql)](https://www.postgresql.org/)
[![Cache](https://img.shields.io/badge/Cache-Redis-DC382D.svg?logo=redis&logoColor=white)](https://redis.io/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A robust, modular, and maintainable RESTful API written in **Go**, powered by the **Gin** web framework, **GORM (PostgreSQL)**, and **Redis**, adhering to **Clean Layered Architecture** principles (`Handler` -> `Service` -> `Repository` / `Client`) organized using an enterprise **Package by Feature (Modular Monolith)** layout.

---

## 🛠️ Tech Stack

- **Language**: [Go](https://go.dev/)
- **Web Framework**: [Gin Web Framework](https://github.com/gin-gonic/gin)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **ORM**: [GORM](https://gorm.io/)
- **In-Memory Cache / Store**: [Redis](https://redis.io/) via [go-redis/v9](https://github.com/redis/go-redis)
- **Configuration**: [godotenv](https://github.com/joho/godotenv)
- **Testing**: Built-in Go `testing` package with interface mocks

---

## 📁 Project Structure

```text
.
├── bootstrap/               # Application bootstrap (router & external client setup)
├── client/                  # Outbound HTTP client integrations
│   ├── city/                # Third-party city service client
│   └── common/              # Common pooled HTTP client, error handlers, & response wrappers
├── internal/                # Private application code
│   ├── config/              # App configurations (App, DB, Redis, Client)
│   └── modules/             # Business feature domains (Package by Feature)
│       ├── city/            # City domain (Handler, Service)
│       ├── products/        # Products domain (Handler, Service, Repository, Entity, Request)
│       ├── role/            # Role domain (Handler, Service, Repository, Entity, Request)
│       └── users/           # Users domain (Handler, Service, Repository, Entity)
├── pkg/                     # Reusable utilities and middlewares (e.g., correlation ID)
├── .env                     # Local environment variables
├── .env.example             # Template environment variables
├── go.mod
├── go.sum
└── main.go                  # Application entry point
```

---

## 🚀 Getting Started

### Prerequisites

Ensure you have the following installed on your machine:
- [Go](https://go.dev/doc/install) (version 1.22+)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Redis](https://redis.io/download/)

### 1. Clone the Repository

```bash
git clone https://github.com/<YOUR_USERNAME>/golang-rest-api.git
cd golang-rest-api
```

### 2. Configure Environment Variables

Create your local `.env` configuration file from the provided `.env.example`:

```bash
cp .env.example .env
```

Sesuaikan konfigurasi environment sesuai kebutuhan Anda:
```env
### Server Config
ENV=development
APP_PORT=8080

### DB Config
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_db

### Redis Config
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
```

### 3. Install Dependencies

Download and verify all project dependencies:

```bash
go mod tidy
```

### 4. Run the Server

Start the application:

```bash
go run main.go
```

The server will initialize the dependencies and start listening on:
```
http://localhost:<APP_PORT> (default: 8080)
```

---

## 🧪 Running Tests

Run all unit tests across the entire codebase:

```bash
go test -v ./...
```

Run tests for a specific module (e.g., `city`):

```bash
go test -v ./internal/modules/city/...
```

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the issues page or submit a pull request.

1. Fork the project
2. Create your feature branch (`git checkout -b feat/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request
