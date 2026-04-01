#!/bin/bash
echo "Checking Docker services..."
docker compose ps
echo ""
echo "Frontend: http://localhost:3000"
echo "Backend: http://localhost:8080"
