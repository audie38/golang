package main

import "fmt"

/**
Struct mirip dengan konsep Class dalam OOP
Biasanya PascalCase
*/

type Customer struct {
	Name, Address string
	Age           int
	Married  bool
}

func main() {
	var ichigo Customer
	ichigo.Name = "Audie Milson"
	ichigo.Address = "Indonesia"
	ichigo.Age = 22
	ichigo.Married = true

	fmt.Println(ichigo)

	// Struct Literals
	Audie := Customer{
		Name: "Milson",
		Address: "Indonesia",
		Age: 22,
		Married: false,
	}

	fmt.Println(Audie)

	Milson := Customer{"Audie", "Indonesia", 22, false}
	fmt.Println(Milson)

	Kurosaki := Customer{"Kurosaki", "Indonesia", 22, false}
	Kurosaki.sayHello()
	Kurosaki.sayHalo("Joko")
	sayHi(Kurosaki, "Joko")
}

//Struct Method, Mirip dengan Function biasa
func(customer Customer)sayHello(){
	fmt.Println("Hello, My Name is", customer.Name)
}

func (customer Customer) sayHalo(name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func sayHi(customer Customer, name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

