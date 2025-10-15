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
│ └── web
│ ├── config
│ │ └── database.go
│ ├── handlers
│ │ ├── auth.go
│ │ └── snippets.go
│ ├── main.go
│ ├── middleware
│ │ └── auth.go
│ ├── renderTemplates
│ │ └── snippets.go
│ ├── routes
│ │ ├── auth.go
│ │ └── snippets.go
│ └── utils
│ └── helpers.go
├── go.mod
├── go.sum
├── internal
│ └── models
│ ├── snippets.go
│ └── user.go
├── README.md
└── ui
├── html
│ ├── pages
│ │ ├── auth
│ │ │ ├── login.html
│ │ │ ├── profile.html
│ │ │ └── register.html
│ │ ├── base.html
│ │ ├── error.html
│ │ └── snippets
│ │ ├── create.html
│ │ ├── home.html
│ │ └── view.html
│ └── partials
│ └── nav.html
└── static
├── css
│ └── main.css
├── img
│ ├── delete.svg
│ ├── favicon.ico
│ └── logo.png
└── js
└── main.js


---

## ⚙️ Setup & Installation

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/Script-Savant/Golang/
cd Golang/snippet-box

2️⃣ Configure the Database

In MySQL, create the database:

CREATE DATABASE snippetbox_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

Then update your DSN (Data Source Name) inside main.go:

dsn := "alex:password@tcp(localhost:3306)/snippetbox_db?parseTime=True&loc=Local"

Replace alex and password with your own MySQL credentials.
3️⃣ Run Database Migrations

GORM automatically migrates tables based on the models in /internal/models.
Just run the app once to auto-create them:

go run ./cmd/web

4️⃣ Run the Application

go run ./cmd/web

Server will start at:

http://localhost:8080

🧪 Sample Data (Optional)

If you want to populate your snippets table with test data:

INSERT INTO snippets (title, content, expires_in, expires_at, user_id)
VALUES
('Welcome to SnippetBox', 'This is your first snippet!', 7, DATE_ADD(NOW(), INTERVAL 7 DAY), 1),
('Gin + GORM Rocks', 'A powerful combination for Go web apps.', 30, DATE_ADD(NOW(), INTERVAL 30 DAY), 2),
('Session Middleware', 'Using secure cookie-based sessions in Gin.', 7, DATE_ADD(NOW(), INTERVAL 7 DAY), 1),
('Dynamic Templates', 'Render multiple template sets dynamically.', 365, DATE_ADD(NOW(), INTERVAL 365 DAY), 2),
('Auto Expiry', 'Snippets automatically expire after the set duration.', 30, DATE_ADD(NOW(), INTERVAL 30 DAY), 1);

🧰 Tech Stack

- Language: Go (1.22+)

- Framework: Gin Web Framework

- ORM: GORM

- Database: MySQL

- Templating: Gin multitemplate

- Sessions: gin-contrib/sessions

- Styling: Vanilla CSS

👨‍💻 Author

Script-Savant
GitHub: https://github.com/Script-Savant
Project: https://github.com/Script-Savant/Golang/tree/main/snippet-box

🧾 License

This project is open-source and available under the MIT License