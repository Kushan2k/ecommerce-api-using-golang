# 🛒 E-commerce API with Golang & Fiber

A high-performance RESTful API for an e-commerce platform built using Go and the Fiber web framework. This project demonstrates how to implement essential e-commerce functionalities such as user authentication, product management, and order processing.

## 🚀 Features

- **User Authentication**: Secure login and registration with JWT-based sessions.
- **Product Management**: CRUD operations for products.
- **In Development**

## 🧰 Technologies Used

- **Go (Golang)**: The programming language used for backend development.
- **Fiber**: A fast and lightweight web framework for Go.
- **MYsql**: models are implemented using GORM orm
- **JWT**: JSON Web Tokens for secure authentication.
- **Golang Dotenv**: For managing environment variables.

## 🛠️ Setup & Installation

### Prerequisites

- Go 1.18 or higher
- MongoDB instance (local or cloud)

### Installation Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/Kushan2k/ecommerce-api-using-golang.git
   cd ecommerce-api-using-golang
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Create a .env file based on the provided .env.sample and set the necessary environment variables
4. Run the application
   ```bash
   go run cmd/main.go
   ```

## Product Structure
```bash
ecommerce-api-using-golang/
├── cmd/                  # Application entry point
├── config/               # Configuration files and environment variables
├── db/                   # Database connection and models
├── middlewares/          # HTTP middlewares
├── models/               # Data models
├── services/             # Business logic and services
├── utils/                # Utility functions
├── .env.sample           # Sample environment variables
├── go.mod                # Go module file
├── go.sum                # Go module checksum file
└── main.go               # Main application file

```

## Testing
You can test the API endpoints using tools like Postman
 or Insomnia Import the provided .rest file for pre-configured requests.
