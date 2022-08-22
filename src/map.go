package main

import "fmt"

func main(){
	//Distinct KeyValuePair
	//deskripsi: map [tipe data key] tipe data value
	person := map[string]string{
		"name" : "Audie",
		"address" : "Jakarta",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["name"])
	delete(person, "address")
	fmt.Println(person)

}