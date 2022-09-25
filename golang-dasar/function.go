package main

import "fmt"

//Type declaration
type Filter func(string)string 
type Blacklist func(string)bool

func main(){
	sayHello()
	sayHelloTo("Audie", "Milson")

	var x int64 = 5
	var y int64 = 2
	jumlah := penjumlahan(x, y)
	fmt.Println(jumlah)

	fmt.Println(getHello("Milson"))

	firstName, lastName := getFullName()
	sayHelloTo(firstName, lastName)

	var kuadrat int64 = 5
	fmt.Println(testingMultiple(firstName, kuadrat))

	fn, mn, ln := getCompleteName()
	fmt.Println(fn, mn, ln)

	var angka = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(jumlahSemua(angka))

	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	
	// Variadic dengan Slice parameter
	fmt.Println(sumAll(angka...))

	// function as Value
	goodbye := getGoodBye
	fmt.Println(goodbye(firstName))

	// function as parameter
	testNama := "Anjing"
	sayHelloWithFilter(testNama, spamFilter)
	testNama = "Kucing"
	sayHelloWithFilter(testNama, spamFilter)

	// Anonymous function
	blacklist := func(name string) bool{
		return name == "Anjing"
	}

	registerUser("Anjing", blacklist)
	registerUser("Ichigo", blacklist)
	registerUser("Ichigo", func(name string)bool{
		return name == "Anjing"
	})
	registerUser("Anjing", func(name string)bool{
		return name == "Anjing"
	})

}

// function
func sayHello(){
	fmt.Println("Hello")
}

// function dengan parameter
func sayHelloTo(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

// function dengan return value
func penjumlahan(first int64, second int64) int64{
	return (first + second)
}

func getHello(name string)string{
	return "Hello " + name
}

// function return multiple value
func getFullName()(string, string){
	return "Audie", "Milson"
}

func testingMultiple(nama string, angka int64)(string, int64){
	return ("Hai " + nama + " ini hasilnya ya: "), (angka*angka)
}

// function dengan Named return values
func getCompleteName() (firstName, middleName, lastName string){
	firstName = "Shigekuni"
	middleName = "Yamamoto"
	lastName = "Genryusai"

	return firstName, middleName, lastName //optional, cukup return
}

// function dengan param slice
func jumlahSemua(numbers []int) int{
	var total int
	total = 0
	for _, val := range numbers{
		total += val
	}

	return total
}

// Variadic Function
func sumAll(numbers ...int)int{ //variadic param harus di posisi param paling akhir
	total:= 0
	for _, number := range numbers{
		total += number
	}

	return total
}

// function Value
func getGoodBye(name string)string{
	return "Good Bye " + name
}

// function as parameter
func spamFilter(name string)string{
	if name == "Anjing"{
		return "..."
	}else{
		return name
	}
}

func sayHelloWithFilter(name string, filter Filter){
	fmt.Println("Hello", filter(name))
}

// Anonymous function
func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You Are Blocked", name)
	}else{
		fmt.Println("Welcome", name)
	}
}
