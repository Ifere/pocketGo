package main

import "fmt"

func greet() string {
	return "hello earth"
}

func main() {
	greeting := greet()
	fmt.Println(greeting)
}
