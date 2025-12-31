//go:build windows

package main

import "fmt"

func getPlatformInfo() string {
	return "Running on Windows"
}

func init() {
	fmt.Println("Windows-specific initialization")
}

