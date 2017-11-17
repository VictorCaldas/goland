package main

import "fmt"

func Sum(x int, y int) int {
	return x + y
}

func getValue() string{
	return "Hello from this another package"
}

func New(name string) (sub * Sub) {
	return &Sub{
		Name: name,
	}
}

func main() {
	fmt.Printf("Hello World! Victor.\n")
	Sum(5, 5)
}