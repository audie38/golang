package golang_goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T){
	totalCpu := runtime.NumCPU()
	totalThread := runtime.GOMAXPROCS(-1)
	totalGoroutine := runtime.NumGoroutine()

	fmt.Println("Total CPU: ", totalCpu)
	fmt.Println("Total Thread: ", totalThread)
	fmt.Println("Total Goroutine: ", totalGoroutine)
}