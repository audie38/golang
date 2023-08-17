package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayName(number int){
	fmt.Println("Display", number)
}

func TestCreateGoroutine(t *testing.T){
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func TestManyGoroutine(t *testing.T){
	for i := 0; i<100000; i++{
		go DisplayName(i)
	}

	time.Sleep(10 * time.Second)
}