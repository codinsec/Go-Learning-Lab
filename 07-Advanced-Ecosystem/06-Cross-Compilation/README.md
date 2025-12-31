# Cross-Compilation

Learn about cross-compilation in Go, building binaries for different platforms from a single machine.

## Learning Objectives

- Understand what cross-compilation is
- Learn to build for different platforms
- Master GOOS and GOARCH environment variables
- Understand CGO limitations with cross-compilation
- Learn build scripts for multiple platforms
- Understand testing cross-compiled binaries

## Concepts Covered

### Basic Cross-Compilation

Build for different platforms using environment variables:

```bash
GOOS=linux GOARCH=amd64 go build
GOOS=windows GOARCH=amd64 go build
GOOS=darwin GOARCH=arm64 go build
```

### Supported Platforms

**Operating Systems:**
- `linux`, `windows`, `darwin` (macOS)
- `freebsd`, `openbsd`, `netbsd`
- `solaris`, `plan9`, `aix`, `dragonfly`

**Architectures:**
- `amd64`, `386`, `arm`, `arm64`
- `mips`, `mips64`, `mips64le`, `mipsle`
- `ppc64`, `ppc64le`, `riscv64`, `s390x`, `wasm`

### Build Scripts

Create scripts to build for multiple platforms:

```bash
#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o app-linux-amd64
GOOS=windows GOARCH=amd64 go build -o app-windows-amd64.exe
GOOS=darwin GOARCH=amd64 go build -o app-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o app-darwin-arm64
```

## Running the Example

```bash
# Navigate to this directory
cd 06-Cross-Compilation

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/07-advanced-ecosystem/cross-compilation

# Run the program
go run main.go

# Build for current platform
go build

# Build for Linux
GOOS=linux GOARCH=amd64 go build

# Build for Windows
GOOS=windows GOARCH=amd64 go build

# Build for macOS (Intel)
GOOS=darwin GOARCH=amd64 go build

# Build for macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v
```

## Key Takeaways

1. **Easy cross-compilation** - Go makes it simple
2. **Environment variables** - Use GOOS and GOARCH
3. **No special tools** - Standard Go toolchain
4. **CGO complexity** - CGO makes it harder
5. **Build scripts** - Automate multi-platform builds
6. **Test on target** - Verify on actual platform

## Common Patterns

**Build for specific platform:**
```bash
GOOS=linux GOARCH=amd64 go build -o myapp
```

**Build with version info:**
```bash
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=1.0.0" -o myapp
```

**Build script:**
```bash
#!/bin/bash
for os in linux windows darwin; do
    for arch in amd64 arm64; do
        GOOS=$os GOARCH=$arch go build -o "dist/app-$os-$arch"
    done
done
```

## Best Practices

- **Use build scripts** - Automate multi-platform builds
- **Test on target** - Verify on actual platform
- **Version info** - Include version in builds
- **CI/CD** - Use CI for cross-compilation
- **Docker** - Use for testing different platforms
- **Document targets** - List supported platforms

## Important Notes

- **No special tools** - Standard Go toolchain works
- **CGO complexity** - Requires cross-compiler with CGO
- **Test on target** - Always test on actual platform
- **File extensions** - Windows needs .exe extension
- **Permissions** - Linux/Unix binaries need execute permission

## CGO and Cross-Compilation

CGO makes cross-compilation more complex:

```bash
# Need cross-compiler
CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm go build

# Or disable CGO
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build
```

## Example Build Script

```bash
#!/bin/bash
set -e

VERSION=$(git describe --tags --always --dirty)
echo "Building version $VERSION"

mkdir -p dist

# Linux
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$VERSION" -o dist/app-linux-amd64
GOOS=linux GOARCH=arm64 go build -ldflags "-X main.version=$VERSION" -o dist/app-linux-arm64

# Windows
GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$VERSION" -o dist/app-windows-amd64.exe

# macOS
GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$VERSION" -o dist/app-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.version=$VERSION" -o dist/app-darwin-arm64

echo "Build complete!"
ls -lh dist/
```

## Testing Cross-Compiled Binaries

1. **On target platform** - Best option
2. **QEMU** - Emulate target platform
3. **Docker** - Use Docker for testing
4. **CI/CD** - Automated testing on multiple platforms
5. **Virtual machines** - Use VMs for testing

## Common Use Cases

- **Release distributions** - Build for all platforms
- **CI/CD pipelines** - Automated builds
- **Embedded systems** - ARM-based devices
- **Cloud deployments** - Different cloud platforms
- **Client applications** - Desktop/mobile apps

## Next Steps

After understanding cross-compilation, proceed to:
- [07-Docker-Deployment](../07-Docker-Deployment/README.md) - Learn about Docker deployment

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

