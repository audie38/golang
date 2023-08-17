package golang_goroutines

import (
	"fmt"
	"strconv"
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

func TestRangeChannel(t *testing.T){
	channel := make(chan string)

	go func(){
		for i := 0; i<10; i++{
			channel <- strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel{
		fmt.Println("Received Data: ", data)
	}

	fmt.Println("Done")
}

func TestSelectChannel(t *testing.T){
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for counter < 2{
		select{
		case data := <- channel1:
			fmt.Println("Data dari channel 1: ", data)
			counter ++
		case data := <-channel2:
			fmt.Println("Data dari channel 2: ", data)
			counter ++
		}
	}
}