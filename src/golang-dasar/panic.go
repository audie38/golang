package main

import "fmt"

// Menghentikan proses aplikasi
func main(){
	runApp(true)
}

func endApp(){
	fmt.Println("End App")
}

func runApp(error bool){
	defer endApp()
	if error{
		panic("ERROR")
	}
}