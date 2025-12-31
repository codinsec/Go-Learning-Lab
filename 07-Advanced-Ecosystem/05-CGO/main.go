package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lmylib

#include <stdio.h>
#include <stdlib.h>
#include "mylib.h"

// Simple C function
int add(int a, int b) {
    return a + b;
}

// C function that allocates memory
char* get_message() {
    char* msg = (char*)malloc(100);
    sprintf(msg, "Hello from C!");
    return msg;
}

// C function that takes a callback
void call_callback(void (*callback)(int)) {
    callback(42);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// This program demonstrates CGO (C bindings in Go)

func main() {
	fmt.Println("=== CGO (C Bindings) ===")
	fmt.Println()

	// 1. Basic C function call
	fmt.Println("1. Basic C Function Call:")
	result := C.add(10, 20)
	fmt.Printf("   C.add(10, 20) = %d\n", result)
	fmt.Println()

	// 2. C string handling
	fmt.Println("2. C String Handling:")
	cMsg := C.get_message()
	defer C.free(unsafe.Pointer(cMsg))
	
	goMsg := C.GoString(cMsg)
	fmt.Printf("   Message from C: %s\n", goMsg)
	fmt.Println()

	// 3. C callback from Go
	fmt.Println("3. C Callback from Go:")
	// Note: This is a simplified example
	// Real callbacks require more complex setup
	fmt.Println("   Callbacks require careful memory management")
	fmt.Println()

	// 4. CGO directives
	fmt.Println("4. CGO Directives:")
	fmt.Println("   //#cgo CFLAGS: -I.        // Include paths")
	fmt.Println("   //#cgo LDFLAGS: -L. -lmylib // Library flags")
	fmt.Println("   //#cgo pkg-config: libssl  // Use pkg-config")
	fmt.Println()

	// 5. Type conversions
	fmt.Println("5. Type Conversions:")
	fmt.Println("   C.int    -> int")
	fmt.Println("   C.double -> float64")
	fmt.Println("   C.char*  -> *C.char -> string (C.GoString)")
	fmt.Println("   string   -> *C.char (C.CString)")
	fmt.Println()

	// 6. Memory management
	fmt.Println("6. Memory Management:")
	fmt.Println("   - C.CString() allocates C memory (must free)")
	fmt.Println("   - C.GoString() converts C string to Go string")
	fmt.Println("   - C.free() to free C-allocated memory")
	fmt.Println("   - Use defer for cleanup")
	fmt.Println()

	// 7. CGO limitations
	fmt.Println("7. CGO Limitations:")
	fmt.Println("   - Slower than pure Go")
	fmt.Println("   - Requires C compiler")
	fmt.Println("   - Cross-compilation complexity")
	fmt.Println("   - Garbage collector pauses")
	fmt.Println("   - Thread overhead")
	fmt.Println()

	// 8. When to use CGO
	fmt.Println("8. When to Use CGO:")
	fmt.Println("   - Interfacing with C libraries")
	fmt.Println("   - Using existing C code")
	fmt.Println("   - Performance-critical C code")
	fmt.Println("   - System-level operations")
	fmt.Println("   - Consider alternatives first")
	fmt.Println()

	// 9. Alternatives to CGO
	fmt.Println("9. Alternatives to CGO:")
	fmt.Println("   - Pure Go implementations")
	fmt.Println("   - cgo-free bindings")
	fmt.Println("   - FFI libraries")
	fmt.Println("   - Separate C service")
	fmt.Println()

	// 10. Key concepts
	fmt.Println("10. Key Concepts:")
	fmt.Println("   - CGO enables C interop")
	fmt.Println("   - Use import \"C\" for C code")
	fmt.Println("   - Careful memory management required")
	fmt.Println("   - Performance overhead exists")
	fmt.Println("   - Use sparingly and carefully")
	fmt.Println()

	// Note: This example won't compile without proper C setup
	// It's for demonstration purposes
	fmt.Println("Note: This example requires C compiler and proper setup.")
	fmt.Println("For a working example, you need:")
	fmt.Println("  1. C compiler (gcc/clang)")
	fmt.Println("  2. C header files")
	fmt.Println("  3. C library files")
	fmt.Println("  4. Proper CGO directives")
}

