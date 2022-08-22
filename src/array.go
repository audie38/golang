package main

import (
	"fmt"
	"strconv"
)

func main(){
	var names [3]string
	names[0] = "Audie"
	names[1] = "Milson"
	names[2] = "Xie"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	names[2] = strconv.Itoa(18) // Convert Int to String
	fmt.Println(names[2])

	var values = [3]int{
		85,
		90,
		95,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	values[1] = 92
	fmt.Println(values)

}