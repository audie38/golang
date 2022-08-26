package main

import "fmt"

type Man struct {
	Name string
}

func main() {
	milson := Man{"Audie"}
	milson.Married()

	fmt.Println(milson.Name)
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}