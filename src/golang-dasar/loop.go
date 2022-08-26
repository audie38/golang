package main

import "fmt"

func main(){
	var counter int64
	counter = 1

	for counter <= 10{
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// For dengan Statement
	for counter := 1; counter <= 10; counter++{
		fmt.Println("Loop ke ", counter)
	}

	// For Range
	slice := []string{"Audie", "Milson"}

	for index, name := range slice{
		fmt.Println("Index ", index, " = ", name)
	}

	// Gunakan _ supaya variabel yang tidak digunakan tidak error
	for _, name := range slice{ 
		fmt.Println(name)
	}

	// Akses MAP
	person := make(map[string]string)
	person["name"] = "Audie"
	person["title"] = "Programmer"

	for key, value := range person{
		fmt.Println(key, "=", value)
	}

}