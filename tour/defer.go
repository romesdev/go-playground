package main

import "fmt"

func main() {
	// executed just after all others actions in function
	defer fmt.Println("world")

	fmt.Println("hello")
}
