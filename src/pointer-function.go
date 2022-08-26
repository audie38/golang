package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	adr := Address{"Jakarta", "DKI Jakarta", ""}
	ChangeCountryToIndonesia(adr)

	fmt.Println(adr) // Value tidak berubah karena pass by value
	var alamat = Address{
		City: "Subang",
		Province: "Jawa Barat",
		Country: "",
	}
	ChangeCountryToIndonesiaPointer(&alamat) //paramnya harusnya pointer
	fmt.Println(alamat)

}

func ChangeCountryToIndonesia(address Address) {
	address.Country = "Indonesia"
}

//paramnya harusnya pointer
func ChangeCountryToIndonesiaPointer(address *Address) {
	address.Country = "Indonesia"
}