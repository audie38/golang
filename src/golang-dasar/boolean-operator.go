package main

import "fmt"

func main(){
	var nilaiAkhir = 80
	var absensi = 20

	var lulusUjian = nilaiAkhir > 60
	var lulusAbsensi = absensi > 17

	fmt.Println(lulusUjian && lulusAbsensi)
}