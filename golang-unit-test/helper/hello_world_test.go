package helper

import "testing"

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Audie")
	if result != "Hello Audie"{
		panic("Result is not Hello Audie")
	}
}