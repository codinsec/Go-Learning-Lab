# 🐹 Go Learning Lab

A comprehensive, structured learning repository for mastering Go (Golang) programming language. This repository provides a systematic, hands-on approach to learning Go from fundamentals to advanced topics, with practical examples and real-world applications.

## 📚 About This Repository

This is a long-term Go Learning Lab designed to take you from a beginner to an advanced Go developer. Each section builds upon the previous one, ensuring a solid foundation before moving to more complex concepts. All code examples follow Go best practices and production standards, making this repository suitable for both learning and reference.

## 🎯 Learning Path

The repository is organized into 7 main sections, each focusing on specific aspects of Go programming:

### 📍 Section 1: Fundamentals
**Location:** `01-Fundamentals/`

Get started with Go! Learn the syntax, basic concepts, and Go's philosophy.

**Topics Covered:**
- Installation & Environment Setup (Go SDK, GOPATH, GOROOT, Go workspace)
- Hello World (First .go file, `go run`, `go build` commands)
- Package Management (`go mod init`, `go.mod`, `go.sum`)
- Variables & Types (`var`, `const`, short variable declaration `:=`, type inference)
- Control Structures (`if`, `else`, `switch`, `for` loop)
- Package Organization (Package visibility, import/export rules)

### 📍 Section 2: Data Structures & Functions
**Location:** `02-Data-Structures-Functions/`

Learn how to organize data and structure your code effectively.

**Topics Covered:**
- Arrays & Slices (Differences, `append`, `copy` operations)
- Maps (Key-value pairs)
- Functions (Multiple return values, named returns)
- Defer, Panic, Recover (Error and flow control mechanisms)
- Pointers (Memory addresses, `&`, `*` operators)

### 📍 Section 3: Structs & Interfaces
**Location:** `03-Structs-Interfaces/`

Go isn't a classic OOP language, but it provides similar structures in a simpler way.

**Topics Covered:**
- Structs (Creating custom data types, struct literals)
- Methods (Receiver functions, value and pointer receivers)
- Interfaces (Duck typing, polymorphism, empty interface `interface{}`)
- Type Assertions & Type Switches (Runtime type checking and conversion)
- Composition (Embedding instead of inheritance)

### 📍 Section 4: Error Handling & Testing
**Location:** `04-Error-Handling-Testing/`

In Go, errors are values. Learn to manage them and test your code effectively.

**Topics Covered:**
- Error Handling (`if err != nil` pattern, custom error types, error wrapping)
- Unit Testing (`testing` package, `go test` command, test coverage)
- Benchmark Tests (Performance measurement with `go test -bench`)
- Table-Driven Tests (Common Go testing pattern)

### 📍 Section 5: Concurrency - Go's Power
**Location:** `05-Concurrency/`

Dive deep into what makes Go truly powerful: concurrency.

**Topics Covered:**
- Goroutines (Lightweight threads, goroutine lifecycle)
- Channels (Safe data communication between goroutines, buffered/unbuffered, directions)
- Select Statement (Managing multiple channel operations, default case)
- Sync Package (WaitGroup, Mutex, RWMutex, Once, Cond, Pool)
- Context (Cancellation signals, timeout management)
- Concurrency Patterns (Worker pools, fan-in/fan-out, pipelines)

### 📍 Section 6: Standard Library & Web
**Location:** `06-Standard-Library-Web/`

Build real-world applications using Go's rich standard library.

**Topics Covered:**
- JSON Processing (`encoding/json`, struct tags, custom marshaling)
- HTTP Server (`net/http` package, routing, middleware patterns)
- HTTP Client (Sending requests, request/response handling)
- File Operations (`os`, `io` packages, file reading/writing, directory operations)
- Database (`database/sql` interface, PostgreSQL/MySQL drivers, connection pooling, transactions)
- Template Engine (`html/template`, `text/template`)
- Logging (`log`, `log/slog` packages, structured logging)
- Configuration (Environment variables, config file parsing, `flag` package)

### 📍 Section 7: Advanced & Ecosystem
**Location:** `07-Advanced-Ecosystem/`

Advanced techniques used in professional Go projects.

**Topics Covered:**
- Generics (Type-safe generic programming, type constraints)
- Reflection (Runtime type analysis with `reflect` package)
- Profiling & Tracing (`pprof` for performance bottlenecks)
- Build Tags (Conditional compilation, platform-specific code)
- CGO (C code integration)
- Cross-Compilation (Building for different platforms)
- Docker & Deployment (Containerization, multi-stage builds)
- Go Tools (`go fmt`, `go vet`, `golangci-lint`, `goimports`, `gopls`)
- Package Management (Dependency versioning, vendoring, private repositories)

## 🚀 Getting Started

### Prerequisites

- **Go 1.18 or later** installed on your system
  - Download from [https://go.dev/dl/](https://go.dev/dl/)
  - Verify installation: `go version`

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/go-learning-lab.git
   cd go-learning-lab
   ```

2. **Navigate to a section:**
   ```bash
   cd 01-Fundamentals
   ```

3. **Each topic folder contains its own Go module:**
   ```bash
   cd 01-Variables
   go mod init github.com/yourusername/go-learning-lab/01-fundamentals/variables
   ```

4. **Run examples:**
   ```bash
   go run main.go
   # or
   go run .
   ```

5. **Run tests:**
   ```bash
   go test
   go test -v          # verbose output
   go test -cover      # with coverage
   go test -bench=.    # run benchmarks
   ```

## 📁 Repository Structure

```
Go-Learning-Lab/
├── 01-Fundamentals/
│   ├── 01-Variables/
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── main_test.go
│   │   └── README.md
│   ├── 02-Control-Structures/
│   │   └── ...
│   └── ...
├── 02-Data-Structures-Functions/
│   └── ...
├── 03-Structs-Interfaces/
│   └── ...
├── 04-Error-Handling-Testing/
│   └── ...
├── 05-Concurrency/
│   └── ...
├── 06-Standard-Library-Web/
│   └── ...
├── 07-Advanced-Ecosystem/
│   └── ...
├── LICENSE
├── roadmap.md
└── README.md
```

## 📖 How to Use This Repository

### For Beginners

1. **Start with Section 1** - Don't skip the fundamentals!
2. **Read each topic's README.md** - Every topic folder contains a detailed README explaining concepts
3. **Run the examples** - Type the code yourself, don't just read
4. **Experiment** - Modify examples, break things, fix them
5. **Write tests** - Practice writing tests for the examples
6. **Move sequentially** - Each section builds on previous knowledge

### For Intermediate Developers

1. **Review fundamentals** - Ensure you understand Go basics
2. **Focus on Section 5** - Concurrency is Go's strength
3. **Practice with Section 6** - Build real applications
4. **Explore Section 7** - Advanced topics for production code

### Best Practices

- **Follow the order** - Sections are designed to be completed sequentially
- **Read the README** - Each topic folder has detailed explanations
- **Run all examples** - Don't just read, execute the code
- **Write your own tests** - Practice is key
- **Use `go fmt`** - Format your code before committing
- **Follow Go idioms** - Check out [Effective Go](https://go.dev/doc/effective_go)

## 🛠️ Go Commands Reference

### Common Commands

```bash
# Initialize a new module
go mod init <module-path>

# Download dependencies
go mod download

# Tidy dependencies
go mod tidy

# Run a program
go run <file.go>
go run .

# Build a program
go build
go build -o <output-name>

# Run tests
go test
go test -v
go test -cover
go test -bench=.

# Format code
go fmt ./...

# Check for errors
go vet ./...

# Install a package
go install <package-path>
```

## 📝 Code Standards

All code in this repository follows:

- **Go formatting standards** (`go fmt`)
- **Effective Go guidelines**
- **Proper error handling** (`if err != nil` pattern)
- **Clear naming conventions** (lowercase package names, exported names start with uppercase)
- **Comprehensive comments** where necessary
- **Test coverage** for all examples

## 🔗 Additional Resources

### Official Resources
- [A Tour of Go](https://go.dev/tour/) - Official interactive tutorial
- [Go by Example](https://gobyexample.com/) - Practical code examples
- [Effective Go](https://go.dev/doc/effective_go) - Go coding standards
- [Go Documentation](https://go.dev/doc/) - Official documentation
- [Go Blog](https://go.dev/blog/) - Articles from the Go team
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) - Code review guide

### Recommended Learning Path

1. Complete [A Tour of Go](https://go.dev/tour/)
2. Work through this repository section by section
3. Read [Effective Go](https://go.dev/doc/effective_go)
4. Practice by building your own projects
5. Contribute to open-source Go projects

## 🤝 Contributing

This is a learning repository. If you find errors, have suggestions, or want to add examples:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/improvement`)
3. Make your changes following Go best practices
4. Commit your changes (`git commit -m 'Add: description'`)
5. Push to the branch (`git push origin feature/improvement`)
6. Open a Pull Request

## 📄 License

This project is licensed under the terms specified in the LICENSE file.

## 🗺️ Roadmap

For a detailed learning roadmap, see [roadmap.md](./roadmap.md).

---

**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

