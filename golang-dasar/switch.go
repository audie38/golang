package main

import "fmt"

func main(){
	name := "Ichigo"

	switch name{
	case "Milson":
		fmt.Println("Hallo Milson")
	case "Audie":
		fmt.Println("Hallo Audie")
	default:
		fmt.Println("Hallo " + name)
	}

	// Switch dengan Short statement
	switch length := len(name); length > 5{
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Switch tanpa kondisi
	size := len(name)

	switch{
	case size > 10:
		fmt.Println("Nama terlalu panjang")
	case size > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar dah")
	}

	

}