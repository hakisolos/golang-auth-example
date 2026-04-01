# Command Line Interface

## Server Entry Point

The `cmd/server` package contains the main entry point for the SparkDB API server.

### main.go

Initializes and runs the Gin web server with all configured routes and middleware.

**Startup Process:**
1. Load environment variables from `.env` file
2. Create Gin router instance
3. Register middleware and routes
4. Start server on port 8080

**Routes Registered:**
- `GET /` - Health check
- `POST /signup` - User registration
- `GET /users` - List all users
- `GET /user/:id` - Get user by ID
- `GET /user/delete/:email` - Delete user by email
- `POST /login` - User login (returns JWT)
- `GET /profile` - Get user profile (protected by AuthMiddleware)

**Environment Variables:**
- `JWT_SECRET` - Secret key for signing JWT tokens (required)

**Running the Server:**

```bash
# Development
go run cmd/server/main.go

# Production (with Docker)
docker run -e JWT_SECRET=your_secret sparkdb-api
```

The server will listen on `http://localhost:8080` by default.

For more details, see [API.md](../API.md) for endpoint documentation.
