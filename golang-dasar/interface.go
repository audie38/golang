package main

import (
	"errors"
	"fmt"
)

/**
Struct yang punya Method yang sama dengan di interface akan otomatis mengimplemen interface
*/

type HasName interface {
	GetName() string
}

type Person struct{
	Name string
}

type Animal struct{
	Name string
}

func main() {
	var Audie Person
	Audie.Name = "Milson"
	sayHallo(Audie)

	cat := Animal{
		Name: "Kocheng Oren",
	}

	sayHallo(cat)

	var data interface{} = Ups(1)
	fmt.Println(data)
	data = Ups(2)
	fmt.Println(data)
	data = Ups(3)
	fmt.Println(data)

	// Interface Error
	var contohError error = errors.New("Err...")
	fmt.Println(contohError.Error())

	hasil, err := Pembagian(100, 0)
	if err == nil{
		fmt.Println("Hasil", hasil)
	}else{
		fmt.Println("Error", err.Error())
	}

}

func sayHallo(hasName HasName) {
	fmt.Println("Halo", hasName.GetName())
}

func (person Person) GetName() string{
	return person.Name
}

func (animal Animal) GetName() string{
	return animal.Name
}

// Interface kosong
func Ups(i int) interface{} {
	if i == 1{
		return 1
	}else if i == 2{
		return true
	}else{
		return "Ups"
	}
}

// Interface Error
func Pembagian (nilai int, pembagi int)(int, error){
	if pembagi == 0{
		return 0, errors.New("Pembagian dengan NOL")
	}else{
		return nilai/pembagi, nil
	}
}