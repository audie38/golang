package helper

import "fmt"

/**

Public Access Modifier = nama function di awali huruf besar
Private Access Modifier = nama function di awali huruf kecil

*/

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string){
	fmt.Println("Goodbye", name)
}