# go-kios

A Go-based API backend for a Kios (Kiosk/Store) application.

## Overview

This project is built using:
- **Go 1.26.3**
- **Chi Router** for HTTP routing
- **PostgreSQL (pgx)** for database interaction
- **JWT** for authentication
- **slog** for logging

## Prerequisites

- Go 1.26 or later
- PostgreSQL database

## Setup & Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yogtanko/go-kios.git
   cd go-kios
   ```

2. **Configure Environment Variables:**
   Copy the example environment file and configure it with your database connection string:
   ```bash
   cp .env.example .env
   ```
   Update the `.env` file with your actual `DBUrl`.

3. **Install Dependencies:**
   ```bash
   go mod tidy
   ```

4. **Run the Application:**
   ```bash
   go run cmd/main.go cmd/api.go
   ```
   The server will start on port `8080` (or as configured).

## Endpoints

- `GET /health` - Healthcheck endpoint
- `POST /login` - User authentication and JWT generation
- `GET /products` - List products (Requires Authentication)

## Project Structure

- `/cmd` - Application entry points (`main.go`, `api.go`)
- `/internal` - Core business logic and internal packages (`auth`, `products`, `postgress`, `middleware`)
- `/pkg` - Shared libraries and configuration (`config`)
