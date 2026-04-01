# Contributing to SparkDB Mock Server

Thank you for your interest in contributing! This project welcomes contributions from everyone.

## Getting Started

1. **Fork** the repository
2. **Clone** your fork:
   ```bash
   git clone https://github.com/yourusername/sparkdb-mock-server.git
   cd sparkdb-mock-server
   ```
3. **Create a branch** for your feature:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Setup

### Backend (Go)

```bash
go mod download
go run cmd/server/main.go
```

### Frontend (Python)

```bash
cd frontend
pip install -r requirements.txt
python app.py
```

### Docker

```bash
docker compose up -d
```

## Making Changes

### Code Style

- **Go**: Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- **Python**: Follow [PEP 8](https://www.python.org/dev/peps/pep-0008/)
- **JavaScript/HTML**: Use standard formatting

### Testing

Before submitting a PR:

1. **Test your changes** locally
2. **Test with Docker** to ensure containers work
3. **Use the frontend** to manually test API endpoints
4. Check HTTP tests in `tests/auth.http`

### Commit Messages

Use clear, descriptive commit messages:

```
Add JWT authentication to profile endpoint
- Implement AuthMiddleware for protected routes
- Convert JWT secret to []byte for correct signing
- Update docker-compose to include environment variables
```

## Submitting a Pull Request

1. **Push** your branch to your fork
2. **Create a PR** with a clear title and description
3. **Link** any related issues
4. **Wait** for review and feedback
5. **Update** based on review comments

### PR Guidelines

- Keep PRs focused on a single feature/fix
- Include relevant documentation updates
- Add API.md updates if endpoints change
- Update docker-compose.yml if dependencies change
- Include clear commit messages

## Project Structure

```
sparkdb-mock-server/
├── cmd/server/          # API server entry point
├── internals/
│   ├── handlers/        # HTTP handlers
│   ├── middlewares/     # Auth middleware
│   ├── models/          # Data structures
│   └── utils/           # Utilities
├── frontend/            # Flask frontend
├── tests/               # HTTP request tests
└── docs/                # Documentation
```

## Areas for Contribution

- 🐛 **Bug Fixes**: Report and fix issues
- ✨ **Features**: Add new endpoints or functionality
- 📚 **Documentation**: Improve docs and examples
- 🧪 **Tests**: Add more test coverage
- 🐳 **Deployment**: Docker/K8s improvements
- 🎨 **Frontend**: UI/UX improvements

## Reporting Issues

Use GitHub Issues to report bugs:

1. **Describe** the problem clearly
2. **Include** steps to reproduce
3. **Attach** error messages or logs
4. **Suggest** a solution if possible

## Questions or Need Help?

- **WhatsApp**: wa.me/2349112171078
- **X (Twitter)**: [@haki_xer](https://x.com/haki_xer)
- Open a discussion on GitHub

## Code of Conduct

Please be respectful and constructive in all interactions. This is an educational project meant to help everyone learn.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

**Happy Contributing!** 🚀
