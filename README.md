# Simple GO Server

A simple Go server that hosts an HTML page. It includes a POST route for user registration, allowing users to register through a form in the `index.html` file. 

This project is part of my journey of learning Go and is not meant to be fully fleshed out.

## Features

- Hosts a static HTML page
- User registration via a POST route
- Simple and easy to understand code structure

## Getting Started

### Prerequisites

- Go version 1.23.6

### Installation

1. Clone the repository:

```bash
    git clone https://github.com/AdrianRRojo/goBackend.git
```
2. CD into directory:
```bash
    cd goBackend
```

3. Run Project:
```bash
    go run .
```

### ENV variables

This project uses the following environment variables:

- `DB`: The name of the database 
- `DB_USER`: The username for the database 
- `DB_HOST`: The host for the database
- `DB_PASS`: The password for the database (optional). If you use a password for your localhost database, you can set it using this variable. Make sure to add the password to the MySQL configuration in `/View/controllers/Register.go`:
```go
Passwd: os.Getenv("DB_PASS"),
```