# Frontend Documentation

## Overview

The frontend is a single-page application built with Python Flask and vanilla JavaScript. It provides a user-friendly interface for interacting with the SparkDB Mock Server API.

## Features

- **Single-Page Routing** - Navigate between pages without page reloads
- **User Authentication** - Sign up, login, and access protected endpoints
- **Black & White Design** - Clean, minimal UI
- **JWT Token Management** - Automatically stores and uses authentication tokens
- **Real-time Feedback** - Shows API responses and errors

## Architecture

### Backend Server (app.py)

A Flask application that serves static files and handles client-side routing.

**Key Features:**
- Static file serving
- Route catch-all for single-page routing (all routes return index.html)
- CORS enabled for backend communication

**Environment:**
- Runs on port 3000
- Flask development mode disabled in production

### Frontend Client (index.html)

A single HTML file containing:
- HTML structure
- CSS styling
- JavaScript logic

**No build process required** - directly open or serve with Flask.

## Pages and Routes

### 1. Home (`/`)
Welcome page with project overview.

### 2. Sign Up (`/signup`)
Register a new user account.
- Username
- Email
- Password

### 3. Login (`/login`)
Authenticate and receive JWT token.
- Email
- Password
- Token automatically saved to localStorage

### 4. Users (`/users`)
View all registered users.
- Displays user cards with ID, username, and email
- No authentication required

### 5. Profile (`/profile`)
View authenticated user's profile.
- Requires valid JWT token
- Shows user's email address
- Token obtained from login

## JavaScript Functions

### Navigation

```javascript
navigateTo(page)
```
Switches between pages without page reloads.

### Authentication

```javascript
handleSignup()
```
Sends POST request to `/signup` with user data.

```javascript
handleLogin()
```
Sends POST request to `/login` and saves token to localStorage.

### API Calls

```javascript
fetchUsers()
```
Retrieves and displays all users.

```javascript
fetchProfile()
```
Retrieves authenticated user profile using stored token.

### Utilities

```javascript
showResponse(elementId, message, type)
```
Displays API response or error in formatted box.

## Local Storage

The frontend uses browser's localStorage to persist JWT tokens:

```javascript
localStorage.getItem('token')    // Retrieve token
localStorage.setItem('token', token)  // Save token
```

This allows users to remain logged in across page refreshes.

## API Integration

**Base URL:** `http://localhost:8080`

All requests use `fetch()` API with CORS enabled.

### Request Headers

```javascript
{
  'Content-Type': 'application/json',
  'Authorization': 'Bearer ' + token  // For protected endpoints
}
```

### Error Handling

All errors are caught and displayed to the user:
- Network errors
- API validation errors
- Authorization errors
- Invalid credentials

## Styling

The frontend uses inline CSS with:
- Black & white color scheme
- Clean typography (Arial)
- Responsive forms
- Hover effects
- Error/success styling

### Color Scheme

- **Background**: White (#ffffff)
- **Text**: Black (#000000)
- **Headers**: Black header with white text
- **Forms**: White inputs with black borders
- **Buttons**: Black background, white text
- **Response boxes**: Light gray background with borders

## Development

### Running Locally

```bash
cd frontend
pip install -r requirements.txt
python app.py
```

Frontend will be available at `http://localhost:3000`

### Docker

```bash
docker build -t sparkdb-frontend ./frontend
docker run -p 3000:3000 sparkdb-frontend
```

### Debugging

Open browser DevTools (F12) to:
- View console errors
- Inspect API requests in Network tab
- Check localStorage contents
- Debug JavaScript

## Future Enhancements

- [ ] Form validation improvements
- [ ] Better error messages
- [ ] Loading indicators
- [ ] Logout functionality
- [ ] User profile editing
- [ ] Dark mode toggle
- [ ] Search/filter users
- [ ] Pagination for users list

## Contributing

See [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines on improving the frontend.

## Contact

For questions or issues with the frontend:
- **WhatsApp**: [wa.me/2349112171078](https://wa.me/2349112171078)
- **X (Twitter)**: [@haki_xer](https://x.com/haki_xer)
