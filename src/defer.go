package main

import "fmt"

func main() {
	runApplication()
}

func logging(){
	fmt.Println("Selesai memanggil function")
}

func runApplication(){
	defer logging() // function bisa dijadwalkan dieksekusi setelah sebuah function selesai dieksekusi walaupun ada error di function
	fmt.Println("Memulai Aplikasi")
}