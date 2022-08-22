package main

import "fmt"

func main(){
	name := "Budi"

	if name == "Audie"{
		fmt.Println("Hello Audie")
	}else if name == "Milson"{
		fmt.Println("Oh! Halo Milson")
	}else{
		fmt.Println("Halo salam kenal " + name)
	}

	// Short Statement
	if length:= len(name); length > 5{
		fmt.Println("Panjang karakter nama terlalu panjang")
	}else{
		fmt.Println("OK!")
	}
}