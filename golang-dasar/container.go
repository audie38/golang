package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// Container list / Double Linked List
	data := list.New()
	data.PushBack("Audie") // Insert ke belakang
	data.PushBack("Milson")
	data.PushBack("Xie")
	data.PushFront("Ichigo") // Insert ke paling awal

	fmt.Println(data.Front().Next().Value) // Ambil data sesudah data paling awal
	fmt.Println(data.Front().Value) // Ambil data paling awal
	fmt.Println(data.Back().Value) // Ambil data terakhir


	fmt.Println("Double LinkedList Depan Sampai Belakang")
	for e := data.Front(); e!= nil; e = e.Next(){
		fmt.Println(e.Value)
	}

	fmt.Println("Double LinkedList Belakang Sampai Depan")
	for e := data.Back(); e!= nil; e = e.Prev(){
		fmt.Println(e.Value)
	}

	// Container Ring / Circular List
	// data2 := ring.New(5)
	var data2 *ring.Ring = ring.New(5)
	for i := 0; i< data2.Len(); i++{
		data2.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data2 = data2.Next()
	}

	data2.Do(func(value interface{}){ // Print isi data Ring
		fmt.Println(value)
	})

	

}