package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
)

func TestContext(t *testing.T){
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T){
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")
	
	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("b"))	
}

func CreateCounter(ctx context.Context) chan int{
	destination := make(chan int)

	go func(){
		defer close(destination)
		counter := 0

		for{
			select {
			case <- ctx.Done():
				return
			default:
				destination  <- counter
				counter ++
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T){
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	destination := CreateCounter(ctx)

	for n := range destination{
		fmt.Println("Counter: ", n)
		if n == 10{
			break
		}
	}

	cancel()
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}