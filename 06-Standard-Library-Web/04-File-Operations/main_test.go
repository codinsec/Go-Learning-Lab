package main

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAndReadFile(t *testing.T) {
	filename := "test_file.txt"
	content := "Test content"

	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
	defer os.Remove(filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	if string(data) != content {
		t.Errorf("Expected %s, got %s", content, string(data))
	}
}

func TestFileInfo(t *testing.T) {
	filename := "test_info.txt"
	content := "Test"
	os.WriteFile(filename, []byte(content), 0644)
	defer os.Remove(filename)

	info, err := os.Stat(filename)
	if err != nil {
		t.Fatalf("Stat failed: %v", err)
	}

	if info.Name() != filename {
		t.Errorf("Expected name %s, got %s", filename, info.Name())
	}

	if info.Size() != int64(len(content)) {
		t.Errorf("Expected size %d, got %d", len(content), info.Size())
	}
}

func TestDirectoryOperations(t *testing.T) {
	dirName := "test_dir"
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		t.Fatalf("Mkdir failed: %v", err)
	}
	defer os.Remove(dirName)

	info, err := os.Stat(dirName)
	if err != nil {
		t.Fatalf("Stat failed: %v", err)
	}

	if !info.IsDir() {
		t.Error("Should be a directory")
	}
}

func TestPathOperations(t *testing.T) {
	path := filepath.Join("dir", "subdir", "file.txt")
	base := filepath.Base(path)
	dir := filepath.Dir(path)
	ext := filepath.Ext(path)

	if base != "file.txt" {
		t.Errorf("Expected 'file.txt', got %s", base)
	}

	if dir != filepath.Join("dir", "subdir") {
		t.Errorf("Expected 'dir/subdir', got %s", dir)
	}

	if ext != ".txt" {
		t.Errorf("Expected '.txt', got %s", ext)
	}
}

func TestFileCopying(t *testing.T) {
	src := "test_src.txt"
	dst := "test_dst.txt"
	content := "Test content"

	os.WriteFile(src, []byte(content), 0644)
	defer os.Remove(src)
	defer os.Remove(dst)

	sourceFile, err := os.Open(src)
	if err != nil {
		t.Fatalf("Open failed: %v", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	defer destFile.Close()

	bytesWritten, err := io.Copy(destFile, sourceFile)
	if err != nil {
		t.Fatalf("Copy failed: %v", err)
	}

	if bytesWritten != int64(len(content)) {
		t.Errorf("Expected %d bytes, got %d", len(content), bytesWritten)
	}

	// Verify copied content
	copiedData, _ := os.ReadFile(dst)
	if string(copiedData) != content {
		t.Errorf("Expected %s, got %s", content, string(copiedData))
	}
}

