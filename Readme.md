# SparkDB Mock Server

A minimal mock API server built with Go and Gin framework for educational purposes. Demonstrates project structure, API design, and deployment with Docker.

**Features:**
- User authentication with JWT
- RESTful API endpoints
- Docker & Docker Compose setup
- Frontend built with Python Flask
- Single-page application routing

## Quick Start

### Prerequisites
- Docker & Docker Compose
- Or: Go 1.25+, Python 3.11+

### Running with Docker Compose

```bash
docker compose up -d
```

Access:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

### Running Locally

**Backend:**
```bash
go mod download
go run cmd/server/main.go
```

**Frontend:**
```bash
cd frontend
pip install -r requirements.txt
python app.py
```

## API Documentation

See [API.md](API.md) for complete endpoint documentation.

## File Structure

```
.
├── cmd/
│   └── server/          # Main server entry point
├── internals/
│   ├── handlers/        # HTTP request handlers
│   ├── middlewares/     # JWT authentication middleware
│   ├── models/          # Data models
│   └── utils/           # Helper functions
├── frontend/            # Python Flask frontend
├── tests/               # HTTP test requests
├── Dockerfile           # Backend container
└── docker-compose.yml   # Multi-service orchestration
```

## Contributing

We welcome contributions! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) for details.

## Contact

- **WhatsApp**: [wa.me/2349112171078](https://wa.me/2349112171078)
- **X (Twitter)**: [@haki_xer](https://x.com/haki_xer)

## Educational Purpose

This repository is for educational purposes teaching:
- Go API development
- Project structure and organization
- Authentication & authorization
- Docker containerization
- Full-stack web applications
- API testing & documentation

