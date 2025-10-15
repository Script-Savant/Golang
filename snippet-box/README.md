# 📝 SnippetBox (Go + Gin + GORM)

SnippetBox is a lightweight web application built with **Go**, **Gin**, and **GORM** that lets users register, log in, and create time-limited text snippets (similar to Pastebin).  
It’s designed for simplicity, security, and clean code organization — ideal as a reference for building modern Go web apps with authentication and database persistence.

---

## 🚀 Features

- 🔐 **User Authentication**
  - Secure registration and login using bcrypt password hashing.
  - Session-based authentication using Gin’s session middleware.
  - Password strength validation (uppercase, lowercase, number, special character).

- 🧾 **Snippet Management**
  - Create snippets with title, content, and expiry duration.
  - Expiry options: **1 day**, **1 week**, or **1 year**.
  - Each snippet is linked to a registered user.
  - Expired snippets can be auto-removed or hidden from listings.

- 🧍 **User Profile**
  - View username and update password securely.
  - Password update requires verifying the old password.

- 💾 **Database**
  - Uses **MySQL** with **GORM ORM**.
  - Includes soft deletion via `deleted_at`.
  - Timestamps for creation and updates managed automatically.

- 🎨 **Templating System**
  - HTML templates rendered via `multitemplate.Renderer`.
  - Shared base layout with dynamic sections for forms and pages.
  - Error messages shown directly within forms.

---

## 🧱 Project Structure

.
├── cmd
│   └── web
│       ├── config
│       │   └── database .go
│       ├── handlers
│       │   ├── auth.go
│       │   └── snippets.go
│       ├── main.go
│       ├── middleware
│       │   └── auth.go
│       ├── renderTemplates
│       │   └── snippets.go
│       ├── routes
│       │   ├── auth.go
│       │   └── snippets.go
│       └── utils
│           └── helpers.go
├── go.mod
├── go.sum
├── internal
│   └── models
│       ├── snippets.go
│       └── user.go
├── README.md
└── ui
    ├── html
    │   ├── pages
    │   │   ├── auth
    │   │   │   ├── login.html
    │   │   │   ├── profile.html
    │   │   │   └── register.html
    │   │   ├── base.html
    │   │   ├── error.html
    │   │   └── snippets
    │   │       ├── create.html
    │   │       ├── home.html
    │   │       └── view.html
    │   └── partials
    │       └── nav.html
    └── static
        ├── css
        │   └── main.css
        ├── img
        │   ├── delete.svg
        │   ├── favicon.ico
        │   └── logo.png
        └── js
            └── main.js



---

## ⚙️ Setup & Installation

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/Script-Savant/Golang/
cd Golang/snippet-box

### 2️⃣ Configure Database
CREATE DATABASE snippetbox_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

dsn := "<your username>:<password>@tcp(localhost:3306)/snippetbox_db?parseTime=True&loc=Local"


## Run the Application
`go run ./cmd/web`
