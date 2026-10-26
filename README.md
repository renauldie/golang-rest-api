# 🚀 Golang REST API

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue.svg?logo=go)](https://go.dev/)
[![Framework](https://img.shields.io/badge/Framework-Gin-00ADD8.svg?logo=gin)](https://github.com/gin-gonic/gin)
[![ORM](https://img.shields.io/badge/ORM-GORM-7952B3.svg)](https://gorm.io/)
[![Database](https://img.shields.io/badge/Database-PostgreSQL-4169E1.svg?logo=postgresql)](https://www.postgresql.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A robust, modular, and maintainable RESTful API written in **Go**, powered by the **Gin** web framework and **GORM (PostgreSQL)**, adhering to **Clean Layered Architecture** principles (`Handler` -> `Service` -> `Repository` / `Client`).

---

## 🛠️ Tech Stack

- **Language**: [Go](https://go.dev/)
- **Web Framework**: [Gin Web Framework](https://github.com/gin-gonic/gin)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **ORM**: [GORM](https://gorm.io/)
- **Configuration**: [godotenv](https://github.com/joho/godotenv)
- **Testing**: Built-in Go `testing` package with interface mocks

---

## 🚀 Getting Started

### Prerequisites

Ensure you have the following installed on your machine:
- [Go](https://go.dev/doc/install) (version 1.22+)
- [PostgreSQL](https://www.postgresql.org/download/)

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

The server will initialize the database schema and start listening on:
```
http://localhost:8080
```

---

## 🧪 Running Tests

Run all unit tests across the entire codebase:

```bash
go test -v ./...
```

Run tests for a specific module (e.g., `city`):

```bash
go test -v ./city/...
```

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the issues page or submit a pull request.

1. Fork the project
2. Create your feature branch (`git checkout -b feat/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request

