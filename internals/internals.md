# Internals Package

The `internals` package contains the core business logic and supporting packages for the API.

## Structure

```
internals/
├── handlers/      # HTTP request handlers
├── middlewares/   # Middleware (authentication, etc.)
├── models/        # Data structures
├── utils/         # Utility functions
└── db.go          # Database/in-memory storage
```

## Main Components

### db.go - In-Memory User Storage

Maintains an in-memory list of users for this mock server.

```go
var Users []models.User
```

All user data is stored in memory and lost when the server restarts.

### handlers/ - HTTP Request Handlers

Processes incoming HTTP requests and returns responses.

**Key Handlers:**
- `CreateUser()` - Handle user registration (POST /signup)
- `GetUsers()` - List all users (GET /users)
- `GetOneUser()` - Get single user (GET /user/:id)
- `DeleteUser()` - Delete user (GET /user/delete/:email)
- `HandleLogin()` - Authenticate user (POST /login)
- `Me()` - Get authenticated user profile (GET /profile)

### middlewares/ - Authentication Middleware

Protects routes that require authentication.

**AuthMiddleware():**
- Extracts Bearer token from Authorization header
- Parses and validates JWT
- Stores user claims in request context
- Aborts request if token is invalid or missing

Protected routes:
- `GET /profile` - Requires valid JWT

### models/ - Data Structures

**User Model:**
```go
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}
```

### utils/ - Utility Functions

Helper functions for common operations.

**Key Functions:**
- `UserExists(email string) bool` - Check if user exists
- `Getuser(id int) (User, bool)` - Retrieve user by ID
- `DeleteUser(email string) error` - Delete user from storage
- `Hashpassword(password string) (string, error)` - Hash password with bcrypt
- `ComparePassword(plain, hashed string) bool` - Verify password

## Authentication Flow

1. User signs up via `/signup` endpoint
   - Password is hashed with bcrypt
   - User added to in-memory storage

2. User logs in via `/login` endpoint
   - Email and hashed password are verified
   - JWT token is generated with 24-hour expiration
   - Token returned to client

3. Client uses token to access protected routes
   - Token sent in `Authorization: Bearer <token>` header
   - Middleware validates signature and expiration
   - Handler accesses user claims from context

## JWT Token Structure

Tokens are signed with HS256 using `JWT_SECRET`.

**Claims:**
- `email` - User's email address
- `exp` - Expiration time (24 hours from login)

Example decoded token:
```json
{
  "email": "user@example.com",
  "exp": 1775162400
}
```

## Error Handling

Handlers return appropriate HTTP status codes:
- `200` - Success
- `400` - Bad request or validation error
- `401` - Unauthorized (invalid token)
- `409` - Conflict (user already exists)
- `500` - Server error

## Development Notes

- No database connection - uses in-memory storage
- Passwords are hashed but not salted per bcrypt (automatic)
- JWTs expire but expired tokens are not actively revoked
- No session management or token blacklists

For client integration, see [API.md](../API.md).
