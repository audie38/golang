package main

import "fmt"

func main(){
	const firstName string = "Audie"
	const lastName = "Milson"
	const value = 1000

	fmt.Println(firstName + " " + lastName)
	fmt.Println(value)

	// Deklarasi multiple constant
	const(
		namaDepan string = "Ichigo"
		namaBelakang = "Kurosaki"
		nilai = 888
	)

	fmt.Println(namaDepan + " " + namaBelakang)
	fmt.Println(nilai)
}