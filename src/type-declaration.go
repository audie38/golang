package main

import "fmt"

func main(){
	type NoKTP string // Membuat alias untuk jenis tipe data
	type Married bool


	var nomorKtp NoKTP = "123123123123123"
	fmt.Println(nomorKtp)
	var marriedStatus Married = false
	fmt.Println(marriedStatus)

}