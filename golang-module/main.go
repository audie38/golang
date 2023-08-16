package main

import (
	"fmt"

	say_hello "github.com/audie38/go-module-sample"
)

func main() {
	fmt.Println(say_hello.SayHello()) 
	say_hello.SayHelloName("Test")
}