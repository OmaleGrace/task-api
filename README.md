# Task API

A RESTful Task Management API built with Go. This project demonstrates the fundamentals of backend development by implementing CRUD (Create, Read, Update, Delete) operations using Go's standard `net/http` package.

The project is part of my journey to mastering backend engineering and learning how REST APIs are built from scratch.

---

## Features

- Create a task
- Retrieve all tasks
- Retrieve a single task by ID
- Update a task
- Delete a task
- JSON request and response handling
- Proper HTTP status codes
- Clean project structure
- Easy to extend with a database and authentication

---

## Tech Stack

- Go
- net/http
- JSON
- REST API Principles

### Planned Technologies

- PostgreSQL
- Docker
- JWT Authentication
- Environment Variables
- Logging
- Unit Testing

---

## API Endpoints

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | `/tasks` | Retrieve all tasks |
| GET | `/tasks/{id}` | Retrieve a task by ID |
| POST | `/tasks` | Create a new task |
| PUT | `/tasks/{id}` | Update an existing task |
| DELETE | `/tasks/{id}` | Delete a task |

---

## Getting Started

### Prerequisites

- Go 1.22 or later

### Clone the Repository

```bash
git clone https://github.com/omalegrace2009-g/task-api.git
cd task-api
```

### Install Dependencies

```bash
go mod tidy
```

### Run the Application

```bash
go run .
```

The server will start at:

```text
http://localhost:8080
```

---

## Example JSON

### Create a Task

**Request**

```json
{
  "title": "Learn REST APIs",
  "completed": false
}
```

**Response**

```json
{
  "id": 1,
  "title": "Learn REST APIs",
  "completed": false
}
```

---

## Learning Objectives

This project is built to understand:

- HTTP fundamentals
- REST API architecture
- Request and response lifecycle
- JSON encoding and decoding
- Routing
- CRUD operations
- Error handling
- Clean project organization
- Backend development with Go

---

## Roadmap

- [x] Initialize Go project
- [x] Create project repository
- [x] Write project documentation
- [ ] Build HTTP server
- [ ] Create API routes
- [ ] Implement CRUD operations
- [ ] Handle JSON requests and responses
- [ ] Add input validation
- [ ] Add PostgreSQL integration
- [ ] Implement JWT authentication
- [ ] Dockerize the application
- [ ] Write unit tests
- [ ] Deploy the API

---

## Future Improvements

- API documentation with Swagger/OpenAPI
- Request logging
- Configuration using environment variables
- Pagination and filtering
- Search functionality
- Rate limiting
- CI/CD with GitHub Actions

---

## Contributing

Contributions, suggestions, and improvements are welcome. Feel free to fork the repository and submit a pull request.

---

## License

This project is licensed under the MIT License.

---

## Author

**Omale Grace**

Backend Developer | Go Enthusiast | Learn2Earn NG Student

GitHub: https://github.com/omalegrace2009-g