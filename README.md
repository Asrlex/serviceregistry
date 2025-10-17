# Go Microservices Service Registry

This is a simple service registry implemented in Go for managing microservices. It allows services to register themselves, discover other services, and perform health checks.

## Features
- Service registration and deregistration
- Service discovery
- Health checks
- RESTful API
- In-memory storage
- Concurrency-safe operations
- Prometheus metrics

## Getting Started

### Prerequisites
- Go 1.16 or higher
- Git

### Installation
1. Clone the repository:
  ```bash
  git clone https://github.com/yourusername/serviceregistry.git
  ```
2. Navigate to the project directory:
  ```bash
  cd serviceregistry
  ```
3. Install dependencies:
  ```bash
  go mod download
  ```
4. Build the project:
  ```bash
  go build -o serviceregistry cmd/main.go
  ```
5. Run the service registry:
  ```bash
  ./serviceregistry
  ```

## Usage
The service registry exposes a RESTful API for service registration, discovery, and health checks. You can interact with the API using tools like `curl` or Postman.
Sample call: 
```bash
curl -X POST http://registry:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "scheduler",
    "url": "http://scheduler:8080",
    "health_url": "http://scheduler:8080/healthcheck"
  }'
```

## Configuration
You can configure the service registry using environment variables. Create a `.env` file in the root