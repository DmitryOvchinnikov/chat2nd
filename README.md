# chat2nd
Study project: chat application with WebSockets, OAuth, etc. 2nd edition

## Structure
```
.
├── README.md
├── auth.go
├── client.go
├── main.go
├── room.go
├── templates
│   └── chat.html
└── trace
    ├── tracer.go
    └── tracer_test.go
```

## Table of Contents
### 1. Chat Application with Web Sockets
- Use the `net/http` package to serve HTTP requests
- Deliver template-driven content to users' browsers
- Satisfy a Go interface to build our own `http.Handler` types
- Use Go's goroutines to allow an application to perform multiple tasks concurrently
- Use channels to share information between running goroutines
- Upgrade HTTP requests to use modern features such as web sockets
- Add tracing to the application to better understand its inner working
- Write a complete Go package using test-driven development practices
- Return unexported types through exported interfaces
### 2. Adding User Accounts
- Use the decorator pattern to wrap `http.Handler` types in order to add additional functionality to handlers
- Serve HTTP endpoints with dynamic paths
- Use the `gomniauth` open source project to access authentication services
- Get and set cookies using the `http` package
- Encode objects as Base64 and back to normal again
- Send and receive JSON data over a web socket
- Give different types of data to templates
- Work with the channels of your own types
### 3. Three Ways to Implement Profile Pictures
### 4. Command-Line Tools to Find Domain Names
### 5. Building Distributed Systems and Working with Flexible Data
### 6. Exposing Data and Functionality through a RESTful Data Web Service API
### 7. Random Recommendations Web Service
### 8. Filesystem Backup
### 9. Building a Q&A Application for Google App Engine
### 10. Micro-services in Go with the Go kit Framework
### 11. Deploying Go Applications Using Docker
