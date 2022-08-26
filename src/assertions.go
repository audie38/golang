package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	switch result := result.(type){
	case string:
		fmt.Println("String", result)
	case int:
		fmt.Println("Int", result)
	default:
		fmt.Println("Unknown")
	}

}