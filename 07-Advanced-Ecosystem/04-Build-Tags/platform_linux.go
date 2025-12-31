//go:build linux

package main

import "fmt"

func getPlatformInfo() string {
	return "Running on Linux"
}

func init() {
	fmt.Println("Linux-specific initialization")
}

