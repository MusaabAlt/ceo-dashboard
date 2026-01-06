# Go Admin

A robust admin panel built with Go, designed for scalability, security, and ease of use.

## Table of Contents
- [Features](#features)
- [Technical Specifications](#technical-specifications)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Features
- User authentication and authorization (JWT-based)
- Role-based access control
- RESTful API architecture
- Modular service structure
- Logging and error handling
- Database integration (PostgreSQL/MySQL)
- Responsive admin dashboard (frontend optional)
- Environment-based configuration

## Technical Specifications

### Backend
- **Language:** Go (>=1.18)
- **Framework:** net/http, gorilla/mux (or similar router)
- **Database:** PostgreSQL (default), MySQL supported
- **Authentication:** JWT tokens
- **ORM:** GORM or sqlx
- **Configuration:** `.env` file or environment variables
- **Logging:** zap or logrus
- **Testing:** go test

### API
- **RESTful endpoints** for CRUD operations
- **JSON** request/response format
- **Pagination** and **filtering** support

### Security
- Password hashing (bcrypt)
- Input validation
- Secure headers (CORS, CSRF protection)
- HTTPS recommended for production

## Installation

```bash
git clone https://github.com/yourusername/go-admin.git
cd go-admin
go mod tidy
```

## Configuration

1. Copy `.env.example` to `.env` and set your environment variables:
    - `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME`
    - `JWT_SECRET`
    - `PORT`

2. (Optional) Edit `config.yaml` for advanced settings.

## Usage

```bash
go run main.go
```
The server will start on the port specified in your `.env` file.

## API Endpoints

| Method | Endpoint         | Description                | Auth Required |
|--------|------------------|----------------------------|--------------|
| POST   | /api/login       | User login                 | No           |
| POST   | /api/register    | User registration          | No           |
| GET    | /api/users       | List users                 | Yes          |
| POST   | /api/users       | Create user                | Yes          |
| PUT    | /api/users/{id}  | Update user                | Yes          |
| DELETE | /api/users/{id}  | Delete user                | Yes          |

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes
4. Push to the branch (`git push origin feature/fooBar`)
5. Open a pull request

## License

MIT License
