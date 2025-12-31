# File Operations

Learn about file and directory operations in Go using `os` and `io` packages.

## Learning Objectives

- Understand how to read and write files
- Learn about file information and metadata
- Master directory operations
- Understand path manipulation with `filepath`
- Learn about file copying and moving
- Understand temporary files and directories

## Concepts Covered

### Reading Files

**Read entire file:**
```go
data, err := os.ReadFile("file.txt")
```

**Read with file handle:**
```go
file, err := os.Open("file.txt")
defer file.Close()
data := make([]byte, 1024)
n, err := file.Read(data)
```

### Writing Files

**Write entire file:**
```go
err := os.WriteFile("file.txt", []byte("content"), 0644)
```

**Write with file handle:**
```go
file, err := os.Create("file.txt")
defer file.Close()
file.WriteString("content")
```

### File Information

```go
info, err := os.Stat("file.txt")
fmt.Println(info.Name())    // File name
fmt.Println(info.Size())   // Size in bytes
fmt.Println(info.Mode())    // File mode
fmt.Println(info.ModTime()) // Modification time
fmt.Println(info.IsDir())   // Is directory
```

### Directory Operations

**Create directory:**
```go
os.Mkdir("dir", 0755)
os.MkdirAll("nested/dir/structure", 0755)
```

**Read directory:**
```go
entries, err := os.ReadDir(".")
for _, entry := range entries {
    fmt.Println(entry.Name())
}
```

**Remove directory:**
```go
os.Remove("dir")
os.RemoveAll("dir") // Recursive
```

### Path Operations

Use `filepath` package for cross-platform paths:

```go
path := filepath.Join("dir", "subdir", "file.txt")
base := filepath.Base(path)      // "file.txt"
dir := filepath.Dir(path)        // "dir/subdir"
ext := filepath.Ext(path)        // ".txt"
clean := filepath.Clean(path)    // Cleaned path
abs, _ := filepath.Abs(path)     // Absolute path
```

### File Copying

```go
src, _ := os.Open("source.txt")
defer src.Close()

dst, _ := os.Create("dest.txt")
defer dst.Close()

io.Copy(dst, src)
```

### Temporary Files

```go
tmpFile, err := os.CreateTemp("", "prefix-*.txt")
defer os.Remove(tmpFile.Name())

tmpDir, err := os.MkdirTemp("", "prefix-*")
defer os.RemoveAll(tmpDir)
```

## Running the Example

```bash
# Navigate to this directory
cd 04-File-Operations

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/06-standard-library-web/file-operations

# Run the program
go run main.go
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover
```

## Key Takeaways

1. **Always close files** - Use defer file.Close()
2. **Use filepath** - For cross-platform paths
3. **Check errors** - File operations can fail
4. **Use ReadFile/WriteFile** - For simple operations
5. **Use file handles** - For more control
6. **Clean up temp files** - Remove temporary files

## Common Patterns

**Read file:**
```go
data, err := os.ReadFile("file.txt")
if err != nil {
    return err
}
```

**Write file:**
```go
err := os.WriteFile("file.txt", data, 0644)
```

**Iterate directory:**
```go
entries, _ := os.ReadDir(".")
for _, entry := range entries {
    // Process entry
}
```

## Best Practices

- **Always close files** - Use defer
- **Use filepath.Join** - For path construction
- **Check errors** - File operations can fail
- **Use appropriate permissions** - 0644 for files, 0755 for dirs
- **Clean up temp files** - Remove when done

## Important Notes

- **File permissions** - Use octal notation (0644, 0755)
- **Cross-platform** - Use filepath, not string concatenation
- **File handles** - Must be closed
- **Temporary files** - Clean up after use

## Next Steps

After understanding file operations, proceed to:
- [05-Database](../05-Database/README.md) - Learn about database operations

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

