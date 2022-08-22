package main

import "fmt"

func main(){
	// Jika tidak tahu kapasitas Array maka bisa gunakan ...
	// jika tidak diisi kapasitasnya maka yang terbentuk adalah slice
	var bulan = [...]string{
		"Januari",
		"Febuari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	//Perubahan data bisa dilakukan pada array/slice 

	var slice1 = bulan[4:7] // mulai dari index 4 sampai 6
	fmt.Println("Slice 1")
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = bulan[4:] // mulai dari index 4
	fmt.Println("Slice 2")
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	var slice3 = bulan[:7] // sampai index 6
	fmt.Println("Slice 3")
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	var slice4 = bulan[:] //semuanya
	fmt.Println("Slice 4")
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	hari := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	fmt.Println(hari)

	var slice5 = hari[:5]
	fmt.Println("Slice 5")
	fmt.Println(slice5)

	var slice6 = append(slice5, "Libur")
	fmt.Println("Slice 6")
	fmt.Println(slice6)

	// Membuat slice dari awal
	newSlice := make([]string, 2, 5) // deskripsi: array tipe data, length, capacity
	newSlice[0] = "Audie"
	newSlice[1] = "Milson"

	fmt.Println(newSlice)

	//Copy Slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice) //to, source
	fmt.Println(copySlice)

}