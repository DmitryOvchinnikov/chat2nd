# chat2nd
Study project: chat application with WebSockets, OAuth, etc. 2nd edition

## Structure
```
.
├── README.md
├── auth.go
├── avatar.go
├── avatar_test.go
├── avatars
│   └── ***.jpg
├── client.go
├── config.go
├── main.go
├── message.go
├── room.go
├── templates
│   ├── chat.html
│   ├── login.html
│   └── upload.html
├── trace
│   ├── tracer.go
│   └── tracer_test.go
└── upload.go
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
- What the good practices to get additional information from auth services are, even when there are no standards in place
- When it is appropriate to build abstractions into our code
- How Go's zero-initialization pattern can save time and memory
- How reusing an interface allows us to work with collections and individual objects in the same way as the existing interface did
- How to use the [https://en.gravatar.com/](https://en.gravatar.com/) web service
- How to do MD5 hashing in Go
- How to upload files over HTTP and store them on a server
- How to serve static files through a Go web server
- How to use unit tests to guide the refactoring of code
- How and when to abstract functionality from struct types into interfaces

### 4. Command-Line Tools to Find Domain Names
- How to build complete command-line applications with as little as a single code file
- How to ensure that the tools we build can be composed with other tools using standard streams
- How to interact with a simple third-party JSON RESTful API
- How to utilize the standard in and out pipes in Go code
- How to read from a streaming source, one line at a time
- How to build a WHOIS client to look up domain information
- How to store and use sensitive or deployment-specific information in environment variables

### 5. Building Distributed Systems and Working with Flexible Data
- Learn about distributed NoSQL datastores, specifically how to interact with MongoDB
- Learn about __distributed messaging queues__, in our case, Bit.ly's NSQ and how to use the `go-nsq` package to easily publish and subscribe to events
- Stream live tweet data through Twitter's streaming APIs and manage long running net connections 
- Learn how to properly stop programs with many internal goroutines 
- Learn how to use low memory channels for signaling

### 6. Exposing Data and Functionality through a RESTful Data Web Service API

### 7. Random Recommendations Web Service

### 8. Filesystem Backup

### 9. Building a Q&A Application for Google App Engine

### 10. Micro-services in Go with the Go kit Framework

### 11. Deploying Go Applications Using Docker
