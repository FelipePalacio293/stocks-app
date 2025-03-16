# Stocks Application

A full-stack application for tracking and analyzing stock recommendations from various brokerages.

## Project Structure

The application consists of two main components:

- `stocks-app-backend/` - Go backend server
- `stocks-app-frontend/` - Vue.js frontend application 

## Backend Features

- RESTful API built with Go and Gin framework
- PostgreSQL database integration using GORM
- Automated stock data synchronization from external API
- Stock recommendations scoring and filtering
- CORS support and environment configuration

## Frontend Features

- Modern UI built with Vue 3 and TypeScript
- Real-time stock recommendations display
- Detailed stock information modals
- Responsive design with Tailwind CSS
- State management with Pinia

## Getting Started

### Backend Setup

1. Navigate to `stocks-app-backend/` directory
2. Create `.env` file with required configuration:
```env
SERVER_PORT=8081
DB_HOST=localhost
DB_PORT=26257
DB_USER=root
DB_PASSWORD=
DB_NAME=stocks
API_BASE_URL=<your-api-url>
API_KEY=<your-api-key>
ALLOWED_ORIGINS=http://localhost:5173