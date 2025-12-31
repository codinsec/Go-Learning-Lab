package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// This program demonstrates file operations in Go

func main() {
	fmt.Println("=== File Operations ===")
	fmt.Println()

	// 1. Reading a file
	fmt.Println("1. Reading a File:")
	readFile()
	fmt.Println()

	// 2. Writing a file
	fmt.Println("2. Writing a File:")
	writeFile()
	fmt.Println()

	// 3. Appending to a file
	fmt.Println("3. Appending to a File:")
	appendFile()
	fmt.Println()

	// 4. File information
	fmt.Println("4. File Information:")
	fileInfo()
	fmt.Println()

	// 5. Directory operations
	fmt.Println("5. Directory Operations:")
	directoryOperations()
	fmt.Println()

	// 6. Path operations
	fmt.Println("6. Path Operations:")
	pathOperations()
	fmt.Println()

	// 7. File copying
	fmt.Println("7. File Copying:")
	fileCopying()
	fmt.Println()

	// 8. Temporary files
	fmt.Println("8. Temporary Files:")
	tempFiles()
	fmt.Println()
}

func readFile() {
	// Method 1: Read entire file
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Printf("   Error reading file: %v\n", err)
		fmt.Println("   (File doesn't exist - this is expected)")
	} else {
		fmt.Printf("   File content: %s\n", string(data))
	}

	// Method 2: Read file in chunks
	file, err := os.Open("example.txt")
	if err == nil {
		defer file.Close()
		buffer := make([]byte, 1024)
		n, _ := file.Read(buffer)
		fmt.Printf("   Read %d bytes\n", n)
	}
}

func writeFile() {
	content := "Hello, Go File Operations!\nThis is a test file."
	
	// Method 1: Write entire file
	err := os.WriteFile("output.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Println("   File written successfully")
	}

	// Method 2: Write with file handle
	file, err := os.Create("output2.txt")
	if err == nil {
		defer file.Close()
		file.WriteString("Line 1\n")
		file.WriteString("Line 2\n")
		fmt.Println("   File written with file handle")
	}
}

func appendFile() {
	file, err := os.OpenFile("append.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		defer file.Close()
		file.WriteString("Appended line\n")
		fmt.Println("   Data appended to file")
	}
}

func fileInfo() {
	info, err := os.Stat("output.txt")
	if err == nil {
		fmt.Printf("   File name: %s\n", info.Name())
		fmt.Printf("   Size: %d bytes\n", info.Size())
		fmt.Printf("   Mode: %v\n", info.Mode())
		fmt.Printf("   Modified: %v\n", info.ModTime())
		fmt.Printf("   Is directory: %t\n", info.IsDir())
	}
}

func directoryOperations() {
	// Create directory
	err := os.Mkdir("testdir", 0755)
	if err == nil {
		fmt.Println("   Directory created: testdir")
	}

	// Create nested directories
	err = os.MkdirAll("nested/dir/structure", 0755)
	if err == nil {
		fmt.Println("   Nested directories created")
	}

	// Read directory
	entries, err := os.ReadDir(".")
	if err == nil {
		fmt.Printf("   Found %d entries in current directory\n", len(entries))
	}

	// Remove directory
	defer os.RemoveAll("testdir")
	defer os.RemoveAll("nested")
}

func pathOperations() {
	// Join paths
	path := filepath.Join("dir", "subdir", "file.txt")
	fmt.Printf("   Joined path: %s\n", path)

	// Base name
	base := filepath.Base(path)
	fmt.Printf("   Base name: %s\n", base)

	// Directory
	dir := filepath.Dir(path)
	fmt.Printf("   Directory: %s\n", dir)

	// Extension
	ext := filepath.Ext(path)
	fmt.Printf("   Extension: %s\n", ext)

	// Clean path
	clean := filepath.Clean(".././dir/file.txt")
	fmt.Printf("   Cleaned path: %s\n", clean)

	// Absolute path
	abs, _ := filepath.Abs(".")
	fmt.Printf("   Absolute path: %s\n", abs)
}

func fileCopying() {
	src := "output.txt"
	dst := "output_copy.txt"

	sourceFile, err := os.Open(src)
	if err == nil {
		defer sourceFile.Close()

		destFile, err := os.Create(dst)
		if err == nil {
			defer destFile.Close()

			bytesWritten, err := io.Copy(destFile, sourceFile)
			if err == nil {
				fmt.Printf("   Copied %d bytes from %s to %s\n", bytesWritten, src, dst)
				os.Remove(dst) // Cleanup
			}
		}
	}
}

func tempFiles() {
	// Create temporary file
	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err == nil {
		defer os.Remove(tmpFile.Name()) // Cleanup
		fmt.Printf("   Temporary file: %s\n", tmpFile.Name())
		tmpFile.WriteString("Temporary content")
		tmpFile.Close()
	}

	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "example-*")
	if err == nil {
		defer os.RemoveAll(tmpDir) // Cleanup
		fmt.Printf("   Temporary directory: %s\n", tmpDir)
	}
}

