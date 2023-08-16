package helper

import "testing"

/*	Go Unit Test Command Line
 *	go test = run all test
 *	go test -v = to show what test functions running
 *	go test -v ./... = to run all test from root folder
 *	go test -v -run=[test functions name] = to run specific test functions
 */

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Audie")
	if result != "Hello Audie"{
		panic("Result is not Hello Audie")
	}
}
func TestHelloWorldDummy(t *testing.T){
	result := HelloWorld("Dummy")
	if result != "Hello Dummy"{
		panic("Result is not Hello Dummy")
	}
}