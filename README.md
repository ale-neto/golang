# Golang


This is a basic user CRUD (Create, Read, Update, Delete) API built in Go, designed to demonstrate clean architecture, layered structure, and proper use of Go with MongoDB and Swagger. The project is based on teachings from the [HunCoding YouTube channel](https://www.youtube.com/@HunCoding) and can be freely used, modified, or studied by others.

> ✅ Inspired and guided by the amazing work of [HunCoding](https://github.com/HunCoding/meu-primeiro-crud-go)


## 📌 Project Information

- **Project Name**: golang
- **Version**: 1.0
- **Host**: http://localhost:8080
- **Repository**: [github.com/ale-neto/golang](https://github.com/ale-neto/golang)

## ✅ Prerequisites

Before getting started, ensure the following tools are installed:

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/ale-neto/golang.git
cd golang
````

### 2. Run the project with Docker Compose

```bash
docker compose up
```

### 3. Or run MongoDB manually and start the app

```bash
docker container run --name golang-mongo -p 27017:27017 -d mongo
go run main.go
```

The app will be running at: [http://localhost:8080](http://localhost:8080)

---

## 🧪 API Testing

Once the application is running, you can test it using Swagger UI:

**[http://localhost:8080/swagger/index.html#](http://localhost:8080/swagger/index.html#)**

Or via tools like `curl` and `Postman`.

### Examples with `curl`:

* **Create user**:

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"name":"John", "email":"john@example.com", "age":30, "password":"Pass@123"}' \
http://localhost:8080/createUser
```

* **Update user**:

```bash
curl -X PUT -H "Content-Type: application/json" \
-d '{"name":"John Doe"}' \
http://localhost:8080/updateUser/{userId}
```

* **Delete user**:

```bash
curl -X DELETE http://localhost:8080/deleteUser/{userId}
```

---

## 🧾 Data Models

### request.UserLogin

```go
email     string // required, must be valid email
password  string // required, min 6 chars, at least one special (!@#$%*)
```

### request.UserRequest

```go
name      string // required, 4–100 chars
email     string // required, valid email
age       int    // required, between 1 and 140
password  string // required, min 6 chars, at least one special
```

### request.UserUpdateRequest

```go
name string // required
age  int    // required
```

### response.UserResponse

```go
id     string
name   string
email  string
age    int
```

### rest\_err.RestErr

```go
code     int
error    string
message  string
causes   []struct {
  field   string
  message string
}
```

---

## 📡 API Endpoints

> For protected endpoints, include the token in the `Authorization` header:
> `Bearer <your-access-token>`

### **POST /createUser**

* Create a new user
* Body: `UserRequest`
* Response: `200 OK`, `400 Bad Request`, `500 Internal Server Error`

### **DELETE /deleteUser/{userId}**

* Delete user by ID
* Path param: `userId`
* Response: `200 OK`, `400`, `500`

### **GET /getUserByEmail/{email}**

* Get user by email
* Path param: `email`
* Response: `200 OK`, `404 Not Found`, `400 Invalid ID`

### **GET /getUserById/{userId}**

* Get user by ID
* Path param: `userId`
* Response: `200 OK`, `404 Not Found`, `400 Invalid ID`

### **POST /login**

* Login with credentials
* Body: `UserLogin`
* Response: `200 OK (returns token)`, `403 Forbidden`

### **PUT /updateUser/{userId}**

* Update user info by ID
* Path param: `userId`
* Body: `UserUpdateRequest`
* Response: `200 OK`, `400`, `500`

---

## 🤝 Contributing

Feel free to fork the project, open issues or submit pull requests.

---

## 🧡 Credits

This project is heavily based on the work of **[HunCoding](https://www.youtube.com/@HunCoding)**.
Special thanks for all the knowledge shared freely on his YouTube channel. This repo exists thanks to that!

---

## 📄 License

This project is open and free to use — no license restrictions apply.

Enjoy the code and happy coding with Go!

```
