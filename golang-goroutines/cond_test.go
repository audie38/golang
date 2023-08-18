package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int){
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T){
	for i := 0; i < 10; i++{
		go WaitCondition(i)
	}

	go func(){
		for i := 0; i < 10; i++{
			cond.Signal() //signal the cond.wait to continue the process
		}
	}()

	// go func(){
	// 	cond.Broadcast() // send the continue process to all goroutine
	// }()

	group.Wait()
}
