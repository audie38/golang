package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Pointer: Pass By Value -> Valuenya yang diduplikasi
	address1 := Address{"Jakarta Barat", "DKI Jakarta", "Indonesia"}
	address2 := address1

	address2.City = "Jakarta Utara"
	fmt.Println("Address 2", address2)
	fmt.Println("Address 1", address1) // Value tidak berubah karena default pass by value
	
	// Pointer
	address3 := &address1
	address3.City = "Jakarta Utara"
	
	fmt.Println("Address 3", address3)
	fmt.Println("Address 1", address1)

	*address3 = Address{"Serpong", "Tanggerang Selatan", "Indonesia"}
	fmt.Println("Address 3", address3)
	fmt.Println("Address 2", address2)
	fmt.Println("Address 1", address1)

	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Indonesia"
	fmt.Println(alamat1)
	fmt.Println(alamat2)


}