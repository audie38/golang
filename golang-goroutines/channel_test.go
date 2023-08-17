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

func GiveMeResponse(channel chan string){
	time.Sleep(2 * time.Second)
	channel <- "Shigekuni Genryusai Yamamoto"
}

func TestChannelAsParameter(t *testing.T){
	chnl := make(chan string)
	defer close(chnl)

	go GiveMeResponse(chnl)

	data := <- chnl
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn (channel chan<- string){
	time.Sleep(2 * time.Second)
	channel <- "Shigekuni Genryusai Yamamoto"
}

func OnlyOut (channel <-chan string){
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T){
	chnl := make(chan string)
	defer close(chnl)

	go OnlyIn(chnl)
	go OnlyOut(chnl)

	time.Sleep(5 * time.Second)
}