//go:build darwin

package main

import "fmt"

func getPlatformInfo() string {
	return "Running on macOS (Darwin)"
}

func init() {
	fmt.Println("macOS-specific initialization")
}

