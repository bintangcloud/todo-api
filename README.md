# Todo API

REST API untuk mengelola Todo dengan sistem authentication dan authorization menggunakan JWT.

Project ini dibuat menggunakan Golang dengan framework Gin, ORM GORM, dan database PostgreSQL.

---

## 🚀 Tech Stack

- Golang
- Gin Framework
- GORM
- PostgreSQL
- JWT Authentication
- bcrypt Password Hashing
- godotenv

---

# ✨ Features

## Authentication

- User Registration
- Password hashing menggunakan bcrypt
- User Login
- Generate JWT Token
- JWT Middleware
- Protected Routes

## User Management

- Get all users
- Update user
- Delete user

## Todo Management

- Create Todo
- Get Todo berdasarkan user yang login
- Update Todo milik sendiri
- Delete Todo milik sendiri

## Security

- Password tidak disimpan dalam bentuk plaintext
- Database credential menggunakan environment variable
- JWT Secret disimpan di `.env`
- User hanya dapat mengakses data miliknya sendiri

---

# 📁 Project Structure

```
todo-api/
│
├── controllers/
│   ├── user_controller.go
│   └── todo_controller.go
│
├── database/
│   └── database.go
│
├── middleware/
│   └── auth_middleware.go
│
├── models/
│   └── models.go
│
├── routes/
│   └── routes.go
│
├── utils/
│   ├── jwt.go
│   └── password.go
│
├── .env
├── .gitignore
├── go.mod
├── main.go
└── README.md
```

---

# ⚙️ Installation

## 1. Clone Repository

```bash
git clone https://github.com/bintangcloud/todo-api.git
```

Masuk ke folder project:

```bash
cd todo-api
```

---

## 2. Install Dependencies

```bash
go mod tidy
```

---

## 3. Setup Environment Variable

Buat file `.env` pada root project.

Contoh:

```env
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=todo_db
DB_PORT=5432

JWT_SECRET=your_secret_key
```

---

## 4. Setup Database

Buat database PostgreSQL:

```sql
CREATE DATABASE todo_db;
```

Project akan melakukan migration otomatis menggunakan GORM.

---

## 5. Running Project

Jalankan:

```bash
go run main.go
```

Server berjalan pada:

```
http://localhost:8080
```

---

# 🔐 Authentication Flow

```
REGISTER

User Input
    |
    ↓
bcrypt Hash Password
    |
    ↓
Save Database


LOGIN

Email + Password
    |
    ↓
Check Password
    |
    ↓
Generate JWT
    |
    ↓
Return Token


REQUEST

JWT Token
    |
    ↓
Middleware Verification
    |
    ↓
Controller
```

---

# 📌 API Documentation

## Public Routes

### Register

```
POST /register
```

Request:

```json
{
    "name": "Bintang",
    "email": "bintang@example.com",
    "password": "12345678"
}
```

---

### Login

```
POST /login
```

Request:

```json
{
    "email": "bintang@example.com",
    "password": "12345678"
}
```

Response:

```json
{
    "token": "jwt_token"
}
```

---

# 🔒 Protected Routes

Semua endpoint berikut membutuhkan JWT.

Header:

```
Authorization: Bearer <token>
```

---

# Todo Endpoint

## Get All Todos

```
GET /todos
```

Response hanya menampilkan Todo milik user yang sedang login.

---

## Create Todo

```
POST /todos
```

Request:

```json
{
    "title": "Belajar Golang Backend"
}
```

User ID otomatis diambil dari JWT.

---

## Update Todo

```
PUT /todos/:id
```

User hanya dapat mengubah Todo miliknya sendiri.

---

## Delete Todo

```
DELETE /todos/:id
```

User hanya dapat menghapus Todo miliknya sendiri.

---

# 🛡️ Security Concept

## Authentication

Menjawab:

> "Siapa user ini?"

Menggunakan:

- JWT Token
- Middleware Verification

---

## Authorization

Menjawab:

> "Apakah user ini boleh melakukan aksi tersebut?"

Implementasi:

- Mengecek `user_id` dari JWT
- Membatasi akses berdasarkan ownership data

---

# 📚 Learning Outcomes

Project ini mempelajari:

- REST API Architecture
- MVC Pattern sederhana
- Gin Routing
- Database Connection
- ORM menggunakan GORM
- Environment Variable
- Password Security
- JWT Authentication
- Middleware
- Authorization

---

# 👩‍💻 Author

Ni Putu Bintang Permatasari