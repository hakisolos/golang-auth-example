# API Documentation

## Overview

SparkDB Mock Server API provides endpoints for user management and authentication.

**Base URL:** `http://localhost:8080`

## Authentication

Protected endpoints require a Bearer token in the Authorization header:

```
Authorization: Bearer <jwt_token>
```

Tokens are obtained via the `/login` endpoint and expire after 24 hours.

## Endpoints

### 1. Health Check

**GET** `/`

Returns API status.

**Response:**
```json
{
  "message": "sparkdb api running"
}
```

---

### 2. User Registration

**POST** `/signup`

Create a new user account.

**Request Body:**
```json
{
  "id": 1,
  "username": "john_doe",
  "email": "john@example.com",
  "password": "secure_password"
}
```

**Response (201):**
```json
{
  "message": "user created successfully",
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "password": "hashed_password"
  }
}
```

**Errors:**
- `400`: Invalid request body
- `409`: User already exists

---

### 3. Login

**POST** `/login`

Authenticate user and receive JWT token.

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "secure_password"
}
```

**Response (200):**
```json
{
  "message": "user logged in successfully",
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Impva25AZXhhbXBsZS5jb20iLCJleHAiOjE3NzUxNjI0MDB9.signature"
}
```

**Errors:**
- `400`: Invalid credentials
- `400`: User does not exist

---

### 4. Get All Users

**GET** `/users`

Retrieve all registered users.

**Response (200):**
```json
[
  {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "password": "hashed_password"
  },
  {
    "id": 2,
    "username": "jane_doe",
    "email": "jane@example.com",
    "password": "hashed_password"
  }
]
```

**Errors:**
- `404`: No users found

---

### 5. Get User by ID

**GET** `/user/:id`

Retrieve a specific user by ID.

**URL Parameters:**
- `id` (integer, required): User ID

**Response (200):**
```json
{
  "id": 1,
  "username": "john_doe",
  "email": "john@example.com",
  "password": "hashed_password"
}
```

**Errors:**
- `500`: Invalid user ID
- `200`: User does not exist (returns message)

---

### 6. Delete User

**GET** `/user/delete/:email`

Delete a user by email address.

**URL Parameters:**
- `email` (string, required): User email

**Response (200):**
```json
{
  "message": "user deleted successfully"
}
```

**Errors:**
- `400`: User does not exist
- `500`: Error deleting user

---

### 7. Get User Profile (Protected)

**GET** `/profile`

**Headers:**
```
Authorization: Bearer <jwt_token>
```

Retrieve the authenticated user's profile.

**Response (200):**
```json
{
  "message": "profile fetched",
  "email": "john@example.com"
}
```

**Errors:**
- `401`: Missing token
- `401`: Invalid token format
- `401`: Invalid token
- `401`: User not found in context

---

## Testing

Use the frontend at `http://localhost:3000` or test with curl:

```bash
# Sign up
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"id":1,"username":"john","email":"john@example.com","password":"pass"}'

# Login
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"pass"}'

# Get profile (replace TOKEN with actual JWT)
curl -X GET http://localhost:8080/profile \
  -H "Authorization: Bearer TOKEN"
```

See [`tests/auth.http`](tests/auth.http) for more examples.

---

## Error Handling

All errors return appropriate HTTP status codes:
- `400`: Bad Request
- `401`: Unauthorized
- `409`: Conflict
- `500`: Internal Server Error

---

## Security Notes

- Passwords are hashed using bcrypt
- JWTs expire after 24 hours
- Keep `JWT_SECRET` environment variable secure
- Always use HTTPS in production
