package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T){
	group := sync.WaitGroup{}
	pool := sync.Pool{
		New: func() interface{}{
			return "New"
		},
	}
	pool.Put("Shigekuni")
	pool.Put("Genryusai")
	pool.Put("Yamamoto")

	for i := 0; i < 10; i++{
		go func(){
			defer group.Done()
			group.Add(1)

			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}
	group.Wait()

	fmt.Println("Done...")
}