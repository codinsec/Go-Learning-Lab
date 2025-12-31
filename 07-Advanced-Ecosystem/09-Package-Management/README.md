# Package Management

Learn about Go module system and package management in Go.

## Learning Objectives

- Understand Go modules (go.mod, go.sum)
- Learn to add and manage dependencies
- Master version management
- Understand vendoring
- Learn about private repositories
- Understand proxy and checksum database

## Concepts Covered

### Module Initialization

Create a new module:

```bash
go mod init <module-path>
```

Example:
```bash
go mod init github.com/user/project
```

### Adding Dependencies

Add dependencies:

```bash
go get <package>
go get <package>@v1.2.3
go get <package>@latest
go get -u <package>  # Update
```

### Managing Dependencies

**Download:**
```bash
go mod download
```

**Tidy (add missing, remove unused):**
```bash
go mod tidy
```

**Vendor:**
```bash
go mod vendor
go build -mod=vendor
```

**Verify:**
```bash
go mod verify
```

## Running the Example

```bash
# Navigate to this directory
cd 09-Package-Management

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/07-advanced-ecosystem/package-management

# Download dependencies
go mod download

# Tidy dependencies
go mod tidy

# Run the program
go run main.go
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v
```

## Key Takeaways

1. **go.mod** - Module definition file
2. **go.sum** - Checksum file for security
3. **go get** - Add/update dependencies
4. **go mod tidy** - Clean dependencies
5. **go mod vendor** - Vendor dependencies
6. **Semantic versioning** - Version management
7. **Module proxy** - Fast dependency downloads

## Common Commands

### Module Management

```bash
# Initialize
go mod init <module-path>

# Add dependency
go get <package>

# Update dependency
go get -u <package>

# Update all
go get -u ./...

# Remove unused
go mod tidy

# Vendor
go mod vendor
```

### Dependency Information

```bash
# List dependencies
go list -m all

# Why dependency
go mod why <module>

# Dependency graph
go mod graph

# Module info
go list -m
```

### Version Management

```bash
# Specific version
go get <package>@v1.2.3

# Latest version
go get <package>@latest

# Pre-release
go get <package>@v1.2.3-beta.1

# Commit hash
go get <package>@abc123
```

## Best Practices

- **Use semantic versioning** - Follow semver
- **Keep go.mod clean** - Use go mod tidy
- **Commit go.sum** - Always commit go.sum
- **Pin versions** - For production
- **Use vendor** - For reproducible builds
- **Private repos** - Configure GOPRIVATE
- **Update regularly** - Keep dependencies updated

## Important Notes

- **go.mod** - Tracks dependencies
- **go.sum** - Ensures integrity (always commit)
- **Semantic versioning** - Major.Minor.Patch
- **Module proxy** - proxy.golang.org
- **Checksum database** - sum.golang.org
- **Vendoring** - For reproducible builds

## Version Management

### Semantic Versioning

- `v1.2.3` - Major.Minor.Patch
- `v0.1.0` - Pre-1.0 versions
- `v1.2.3-beta.1` - Pre-release
- `v1.2.3+metadata` - Build metadata

### Version Selection

- `@latest` - Latest version
- `@v1.2.3` - Specific version
- `@v1.2` - Latest v1.2.x
- `@master` - Branch/tag name

## Private Repositories

### Configuration

**GOPRIVATE:**
```bash
go env -w GOPRIVATE=github.com/company/*
```

**Authentication:**
- Use `.netrc` or `.gitconfig`
- Use personal access tokens
- Use SSH keys

**Disable checksum:**
```bash
go env -w GOSUMDB=off
```

## Module Proxy

### Default Settings

```bash
GOPROXY=https://proxy.golang.org,direct
GOSUMDB=sum.golang.org
```

### Custom Proxy

```bash
go env -w GOPROXY=https://proxy.company.com,direct
```

### Direct Mode

```bash
go env -w GOPROXY=direct
```

## Vendoring

### Create Vendor Directory

```bash
go mod vendor
```

### Use Vendored Dependencies

```bash
go build -mod=vendor
go test -mod=vendor
```

### Benefits

- Reproducible builds
- Offline builds
- Version control
- Faster CI/CD

## go.mod File Structure

```go
module github.com/user/project

go 1.21

require (
    github.com/pkg/errors v0.9.1
    golang.org/x/sync v0.1.0
)

replace (
    old/module => new/module v1.0.0
)

exclude (
    bad/module v1.0.0
)
```

## Common Patterns

**Add dependency:**
```bash
go get github.com/pkg/errors
```

**Update dependency:**
```bash
go get -u github.com/pkg/errors
```

**Remove unused:**
```bash
go mod tidy
```

**Vendor:**
```bash
go mod vendor
```

**Verify:**
```bash
go mod verify
```

## Troubleshooting

**Clear module cache:**
```bash
go clean -modcache
```

**Reset go.sum:**
```bash
go mod tidy
```

**Check module info:**
```bash
go list -m -versions <module>
```

## Next Steps

Congratulations! You've completed all sections of the Go Learning Lab.

Review what you've learned:
- [01-Fundamentals](../../01-Fundamentals/README.md)
- [02-Data-Structures-Functions](../../02-Data-Structures-Functions/README.md)
- [03-Structs-Interfaces](../../03-Structs-Interfaces/README.md)
- [04-Error-Handling-Testing](../../04-Error-Handling-Testing/README.md)
- [05-Concurrency](../../05-Concurrency/README.md)
- [06-Standard-Library-Web](../../06-Standard-Library-Web/README.md)
- [07-Advanced-Ecosystem](../README.md)

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

