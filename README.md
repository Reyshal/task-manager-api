# 🧠 Task Manager API

A lightweight REST API built with [Fiber](https://gofiber.io/) and [GORM](https://gorm.io/) using Go — built for a simple task management app with authentication using JWT.

## ⚙️ Tech Stack

- [Go](https://golang.org/) 1.21+
- [Fiber](https://docs.gofiber.io/) — web framework
- [GORM](https://gorm.io/) — ORM for database
- [JWT](https://docs.gofiber.io/api/middleware/jwt) — authentication middleware
- PostgreSQL / SQLite

---

## 🚀 Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/Reyshal/task-manager-api.git
cd task-manager-api
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Set up `.env` file
Copy and rename `.env.example` into a `.env` file at root, you can modified it as you like:
```env
PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/taskdb
JWT_SECRET=your_jwt_secret_key
```

### 4. Run the app
```bash
go run main.go
```

---

## 📦 API Endpoints

### 🔐 Auth
- `POST /api/register` — create new user
- `POST /api/login` — get JWT token

### ✅ Tasks (Protected with JWT)
- `GET /api/tasks` — list tasks
- `POST /api/tasks` — create new task
- `PUT /api/tasks/:id` — update task
- `DELETE /api/tasks/:id` — delete task

Use `Authorization: Bearer <token>` header for protected routes.

---

## 🗂 Project Structure

```bash
.
├── main.go
├── database/
│   └── db.go
├── models/
│   └── user.go
│   └── task.go
├── handlers/
│   └── auth.go
│   └── task.go
├── middleware/
│   └── jwt.go
├── routes/
│   └── routes.go
└── .env
```

---

## 🧪 Todo / Features

- [ ] JWT Auth
- [ ] Password hashing with bcrypt
- [ ] Unit tests
- [ ] Dockerfile

---

## 📄 License

MIT — free to use & modify.
