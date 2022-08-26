package main

import "fmt"

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