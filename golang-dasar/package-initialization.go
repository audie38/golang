package main

import (
	"fmt"
	"golang-dasar/database"
)

//import _ "golang-dasar/database" //blank identifier, agar jika import tidak digunakan tidak menyebabkan error

func main(){
	result := database.GetDatabase()
	fmt.Println(result)
}