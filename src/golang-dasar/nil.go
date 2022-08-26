package main

import "fmt"

/**
Hanya bisa digunakan pada interface, function, map, slice, pointer, dan channel
selain tipe data di atas, tidak bisa menggunakan nil jadi menggunakan default valuenya
*/

func main() {
	var person map[string]string = nil

	fmt.Println(person)

	data := NewMap("")
	fmt.Println(data)
}

func NewMap(name string)map[string]string{
	if name == ""{
		return nil
	}else{
		return map[string]string{
			"name": name,
		}
	}
}