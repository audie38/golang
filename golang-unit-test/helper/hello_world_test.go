package helper

import "testing"

/*	Go Unit Test Command Line
 *	go test = run all test
 *	go test -v = to show what test functions running
 *	go test -v ./... = to run all test from root folder
 *	go test -v -run=[test functions name] = to run specific test functions
 */

/*
 * t.Fail() = failed the unit test and continue
 * t.FailNow() = failed the unit test and not continue
 * t.Error() = similar to t.Fail() but can pass failed feedback arguments
 * t.Fatal() = similar to t.FailNow() but can pass failed feedback arguments
 */

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Audie")
	if result != "Hello Dummy"{
		t.Error("Result must be Hello Dummy")
	}
}
func TestHelloWorldDummy(t *testing.T){
	result := HelloWorld("Dummy")
	if result != "Hello test"{
		t.Fatal("Result must be Hello test")
	}
}