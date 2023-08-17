package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ichigo Kurosaki"
		fmt.Println("Finish sending data to channel")
	}()

	data := <- channel
	fmt.Println(data)
	
	time.Sleep(5 * time.Second)
}