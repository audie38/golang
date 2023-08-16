package main

import (
	"errors"
	"fmt"
	"strconv"
)

func loggingApp(){
	message := recover()
	if message != nil{
		fmt.Println("Logging Started...")
		fmt.Println(message)
		fmt.Println("Logging Finished...")
	}
}

func plus(a, b int) (result int) {
	result = a + b
	return
}

func minus(a, b int) (result int) {
	result = a - b
	return
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func divide(a, b float64) (result float64) {
	if b == float64(0){
		panic("Error: Divide by zero error")
	}
	
	result = a / b
	return
}

func divisionWithCustomErrors(a, b float64) (float64, error){
	if b == 0{
		return 0, errors.New("Invalid Divider")
	}
	return a/b, nil
}

func modulo(a, b int) (result int) {
	result = a % b
	return
}

func sumArrayEl(slice ...int) (total int){
	total = 0
	for i := 0; i < len(slice); i++{
		total += slice[i]
	}
	return 
}

type HomeAddress struct{
	City, Province, Country string
}

type UserProfile struct{
	Name, BirthPlace string
	Age int
	Address HomeAddress
}

func SetDefaultHomeAddress(addr *HomeAddress){
	addr.City = "Rukongai"
	addr.Province = "District 38"
	addr.Country = "Soul Society"
}

func SetDefaultHomeAddressNoPointer(addr HomeAddress){
	addr.City = "Rukongai"
	addr.Province = "District 38"
	addr.Country = "Soul Society"
}

type HasInfo interface{
	GetBio() string
	GetAge() int
	GetAddress() string
}

func DisplayUserAddress(hasInfo HasInfo){
	fmt.Println("User's Address: ", hasInfo.GetAddress())
}

func DisplayUserAge(hasInfo HasInfo){
	fmt.Println("User's age : ", hasInfo.GetAge())
}

func DisplayUserBio(hasInfo HasInfo){
	fmt.Println("User's Bio: ", hasInfo.GetBio())
}

func (user *UserProfile) GetBio() string{
	return user.Name + ", is " + strconv.Itoa(user.Age) + " years old that come from " + user.BirthPlace 
}

func (user *UserProfile) GetAge() int{
	return user.Age
}

func (user *UserProfile) GetAddress() string{
	return user.Address.City + ", " + user.Address.Province + ", " + user.Address.Country
}

func main() {
	defer loggingApp()

	var numberArray = []int{1, 2, 3, 4, 5}
	numberArrayTotal := sumArrayEl(numberArray...)
	fmt.Println("Total: ", numberArrayTotal)

	var firstInput int = 10
	var secondInput int = 15

	divisionResult, err := divisionWithCustomErrors(float64(firstInput), float64(secondInput))
	if err != nil{
		fmt.Println("Custom Error: ", err.Error())
	}else{
		fmt.Println("Division with custom error result: ", divisionResult)
	}

	plusResult := plus(firstInput, secondInput)
	minusResult := minus(firstInput, secondInput)
	multiplyResult := multiply(firstInput, secondInput)
	divideResult := divide(float64(firstInput), float64(secondInput))
	moduloResult := modulo(firstInput, secondInput)
	
	fmt.Println("Result : ", plusResult, minusResult, multiplyResult, divideResult, moduloResult)

	var person1 *UserProfile = new(UserProfile)
	person1.Name = "user one"
	person1.Age = 18
	person1.BirthPlace = "Soul Society"
	
	fmt.Println("Before: ", *person1)

	address1 := &HomeAddress{
		City: "",
		Province: "",
		Country: "",
	}

	SetDefaultHomeAddress(address1)
	fmt.Println("Default Home Address: ", *address1)

	person1.Address = *address1
	fmt.Println("After: ", *person1)
	DisplayUserBio(person1)
	DisplayUserAge(person1)
	DisplayUserAddress(person1)

	addressNoPointer := HomeAddress{
		City: "",
		Province: "",
		Country: "",
	}

	SetDefaultHomeAddressNoPointer(addressNoPointer)
	fmt.Println("Address No Pointer: ", addressNoPointer)
}