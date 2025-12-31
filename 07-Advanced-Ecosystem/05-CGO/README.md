# CGO

Learn about CGO, which enables calling C code from Go programs.

## Learning Objectives

- Understand what CGO is and when to use it
- Learn CGO syntax and directives
- Master C type conversions
- Understand memory management with CGO
- Learn CGO limitations and alternatives
- Understand best practices

## Concepts Covered

### CGO Basics

CGO allows calling C code from Go:

```go
/*
#include <stdio.h>
int add(int a, int b) {
    return a + b;
}
*/
import "C"

func main() {
    result := C.add(10, 20)
}
```

### CGO Directives

Control CGO compilation:

```go
//#cgo CFLAGS: -I.        // Include paths
//#cgo LDFLAGS: -L. -lmylib // Library flags
//#cgo pkg-config: libssl  // Use pkg-config
```

### Type Conversions

Convert between C and Go types:

```go
// Go to C
cStr := C.CString("hello")
defer C.free(unsafe.Pointer(cStr))

// C to Go
goStr := C.GoString(cStr)
```

### Memory Management

Carefully manage C-allocated memory:

```go
cStr := C.CString("hello")
defer C.free(unsafe.Pointer(cStr))  // Always free!
```

## Running the Example

```bash
# Navigate to this directory
cd 05-CGO

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/07-advanced-ecosystem/cgo

# Note: CGO requires C compiler (gcc/clang)
# Enable CGO (default on most systems)
export CGO_ENABLED=1

# Run the program
go run main.go

# Build
go build
```

## Running Tests

```bash
# Run tests (requires CGO)
CGO_ENABLED=1 go test

# Disable CGO for tests
CGO_ENABLED=0 go test
```

## Key Takeaways

1. **C interop** - Call C code from Go
2. **Type conversions** - Convert between C and Go types
3. **Memory management** - Must free C-allocated memory
4. **Performance overhead** - CGO has overhead
5. **Cross-compilation** - More complex with CGO
6. **Use alternatives** - Prefer pure Go when possible

## Common Patterns

**Calling C functions:**
```go
/*
int add(int a, int b) {
    return a + b;
}
*/
import "C"

result := C.add(10, 20)
```

**String handling:**
```go
cStr := C.CString("hello")
defer C.free(unsafe.Pointer(cStr))
goStr := C.GoString(cStr)
```

**Using C libraries:**
```go
//#cgo LDFLAGS: -lmylib
//#include "mylib.h"
import "C"
```

## Best Practices

- **Use sparingly** - CGO has overhead
- **Free memory** - Always free C-allocated memory
- **Use defer** - Defer C.free() calls
- **Type safety** - Be careful with type conversions
- **Error handling** - Check C return values
- **Consider alternatives** - Pure Go is often better

## Important Notes

- **C compiler required** - Need gcc or clang
- **Performance overhead** - Slower than pure Go
- **Garbage collector** - GC pauses affect C code
- **Cross-compilation** - More complex with CGO
- **Thread overhead** - CGO uses threads
- **CGO_ENABLED** - Can disable CGO with env var

## CGO Limitations

1. **Performance** - Overhead from Go-C boundary
2. **Compilation** - Requires C compiler
3. **Cross-compilation** - More complex
4. **GC pauses** - Can affect C code
5. **Thread overhead** - Uses threads for C calls
6. **Build complexity** - More complex builds

## When to Use CGO

- **Existing C libraries** - When you need C libraries
- **System calls** - Low-level system operations
- **Performance-critical** - When C is faster
- **Legacy code** - Integrating existing C code
- **No alternative** - When pure Go isn't possible

## Alternatives to CGO

1. **Pure Go** - Rewrite in Go
2. **cgo-free bindings** - Use existing bindings
3. **FFI libraries** - Foreign function interface
4. **Separate service** - Run C code separately
5. **Go libraries** - Use Go equivalents

## Example: Using C Library

```go
//#cgo LDFLAGS: -lmylib
//#include "mylib.h"
import "C"

func main() {
    result := C.my_function(C.int(42))
}
```

## Enabling/Disabling CGO

```bash
# Enable CGO (default)
export CGO_ENABLED=1

# Disable CGO
export CGO_ENABLED=0

# Check CGO status
go env CGO_ENABLED
```

## Next Steps

After understanding CGO, proceed to:
- [06-Cross-Compilation](../06-Cross-Compilation/README.md) - Learn about cross-compilation

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

