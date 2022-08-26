package main

import "fmt"

// Menangkap data panic, panic tidak menghentikan aplikasi
func main() {
	runApp(true)
}

func runApp(error bool){
	defer endApp()

	if error{
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi Berjalan")
}

func endApp(){
	message := recover() //disimpan di bagian defer function
	if message != nil{
		fmt.Println("Err message: ", message)
	}else{
		fmt.Println("Aplikasi Selesai")
	}
}