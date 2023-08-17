package helper

import "testing"

/*	Go Unit Test Command Line
 *	go test = run all test
 *	go test -v = to show what test functions running
 *	go test -v ./... = to run all test from root folder
 *	go test -v -run=[test functions name] = to run specific test functions
 */

/*
 * t.FailNow() = failed the unit test and not continue
 */

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Audie")
	if result != "Hello Dummy"{
		t.Fail()
	}
}
func TestHelloWorldDummy(t *testing.T){
	result := HelloWorld("Dummy")
	if result != "Hello test"{
		t.FailNow()
	}
}