# Event Sourcing Golang Implementation

This is a Golang implementation inspired by the Kotlin/Spring Boot event sourcing example, using MySQL as the database.

## Architecture

The project follows a similar structure to the original Kotlin implementation with:

- **Common interfaces**: Event, Command, Query, Processor, ReadModel
- **Domain models**: Event sourcing patterns
- **Infrastructure**: MySQL database with GORM

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.21+ (for local development)

### Running with Docker Compose

1. Start the services:
```bash
docker-compose up -d
```

2. The application will be available at `http://localhost:8080`
3. MySQL will be available at `localhost:3306`

### Running Locally

1. Start MySQL:
```bash
docker-compose up mysql -d
```

2. Run the application:
```bash
go run main.go
```

## Database

- **Database**: eventsourcing
- **Username**: app_user
- **Password**: app_password
- **Root Password**: password

The database is automatically initialized with tables for event store and read models.

## API Endpoints

- `GET /health` - Health check endpoint

## Project Structure

```
├── internal/
│   ├── common/          # Common interfaces and types
│   ├── domain/          # Domain models
│   └── infrastructure/  # Database and external services
├── docker/
│   └── mysql/
│       └── init/        # MySQL initialization scripts
├── docker-compose.yaml
├── Dockerfile
└── main.go
```