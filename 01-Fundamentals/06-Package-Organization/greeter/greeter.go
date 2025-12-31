package greeter

import "fmt"

// SayHello is an exported function
func SayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// SayGoodbye is an exported function
func SayGoodbye(name string) {
	fmt.Printf("Goodbye, %s!\n", name)
}

// sayHello is an unexported function (not accessible outside package)
func sayHello(name string) {
	fmt.Printf("hello, %s\n", name)
}

