package main

import "fmt"

func main(){
	var name string

	name = "Audie Milson"
	fmt.Println(name)

	name = "Ichigo Kurosaki"
	fmt.Println(name)

	var nama = "Edison"
	fmt.Println(nama)

	var age int
	
	age = 22
	fmt.Println(age)

	var umur = 21 // Membuat variable tanpa tipe datanya bisa jika langsung diberi isinya
	fmt.Println(umur)

	var usia int8 = 20 // Override default tipe data
	fmt.Println(usia)

	country := "Indonesia" // Cara membuat variable tanpa var
	fmt.Println(country)

	country = "Singapore"
	fmt.Println(country)

	// Cara untuk membuat deklarasi multiple variabel 
	var (
		firstName = "Audie"
		lastName = "Milson"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	fmt.Println(firstName + " " + lastName)

}