package main

import "fmt"
var greeting = "Hello, Naza!"


func main() {
	fmt.Println(greeting)
}

func shareWith(name string) string {
	if name=="" {
		name = "you"
	}
	return "One for " + name + ", one for me"
}