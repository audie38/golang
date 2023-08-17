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

func TestBufferedChannel(t *testing.T){
	//by default channel can only send & receive 1 (no buffer), to send/receive more than 1 must add buffer
	channel := make(chan string, 3) 
	defer close(channel)

	go func(){
		channel <- "Shigekuni"
		channel <- "Genryusai"
		channel <- "Yamamoto"
	}()

	go func(){
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}