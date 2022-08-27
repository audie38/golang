package main

import (
	"container/list"
	"fmt"
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

}