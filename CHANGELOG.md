# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial project setup with Go and Gin framework
- JWT authentication with 24-hour expiration
- User management endpoints (signup, login, profile)
- Python Flask frontend with single-page routing
- Docker and Docker Compose configuration
- Comprehensive API documentation
- Contributing guidelines and code of conduct
- MIT License

### Fixed
- JWT token signing to use []byte instead of string

### Security
- Password hashing with bcrypt
- JWT token validation in AuthMiddleware
- Environment variable support for JWT_SECRET

---

## Future Planned Features

- [ ] Database integration (PostgreSQL/MongoDB)
- [ ] Email verification
- [ ] Password reset functionality
- [ ] Rate limiting
- [ ] Input validation improvements
- [ ] Comprehensive test suite
- [ ] API versioning
- [ ] Swagger/OpenAPI documentation
- [ ] Refresh token mechanism
- [ ] Role-based access control (RBAC)
