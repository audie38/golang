package main

import (
	"container/list"
	"container/ring"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Filter func(string)string


type Customer struct{
	Name, Address string
	Age int
}

type Person struct{
	Name string
}

// Sort Algo
type User struct{
	Name string
	Age int
}

type UserSlice []User

func (value UserSlice) Len() int{
	return len(value)
}

func (value UserSlice) Less(i, j int) bool{
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int){
	value[i], value[j] = value[j], value[i]
}

// End of Sort Algo interface functions

type Address struct{
	City, Province, Country string
}

type HasName interface{
	GetName() string
}

func (person Person) GetName() string{
	return person.Name
}

func SayHello(hasName HasName){
	fmt.Println("Hello", hasName.GetName())
}

func random() interface{}{
	return "OK"
}

func spamFilter(text string) string{
	if text == "test"{
		return "..."
	}else{
		return text
	}
}

func sayHelloWithFilter(name string, filter Filter){
	fmt.Println("Hello", filter(name))
}

// named return values
func getCompleteName() (firstName, lastName string){
	firstName = "Ichigo"
	lastName = "Kurosaki"
	return
}

// variadic function
func sumAll(numbers ...int) (total int){
	total = 0

	for _, number := range numbers{
		total += number
	}

	return
}

func newMap(name string) map[string]string{
	if name == ""{
		return nil
	}

	return map[string]string{
		"name": name,
	}
}

func (cust Customer) greetCustomer(){
	fmt.Println("Welcome ", cust.Name)
}

func division(val int, div int) (int, error){
	if div == 0{
		return 0, errors.New("Divided with 0")
	}

	return val/div, nil
}

func logging(){
	fmt.Println("Logging Function")
}

func runApplication(value int){
	defer logging()
	fmt.Println("Run App")
	result:= 10/value
	fmt.Println("Result: ", result)
}

func endApp(){
	fmt.Println("End App")
	message := recover()
	if message != nil{
		fmt.Println("Error : ", message)
	}
}

func runAppPanic(err bool){
	defer endApp()
	if err{
		panic("ERROR")
	}
}

func ChangeAddressToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func main() {
	var newSlice = make([]int, 2, 5)
	newSlice[0] = 1
	newSlice[1] = 2
	var newArray = [...]int{1, 2, 3, 4, 5, 6}
	
	fmt.Println(newSlice, newArray)

	copySlice := make([]int, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	sliceFromArr := newArray[3:6] // index start from : index end before
	fmt.Println(sliceFromArr)
	fmt.Println("Slice Len: ", len(sliceFromArr))
	fmt.Println("Slice Cap: ", cap(sliceFromArr))

	slice2 := append(sliceFromArr, 7)
	fmt.Println(slice2)
	fmt.Println("Slice Len: ", len(slice2))
	fmt.Println("Slice Cap: ", cap(slice2))

	var fullName string = "@Ichigo Kurosaki"
	var days string = "Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday"
	
	fmt.Println(strings.Trim(fullName, "@"))
	fmt.Println(strings.ToLower(fullName))
	fmt.Println(strings.ToUpper(fullName))
	fmt.Println(strings.Split(days, ","))
	fmt.Println(strings.Contains(days, "Saturday"))
	fmt.Println(strings.Contains(days, "Holiday"))
	fmt.Println(strings.ReplaceAll(days, "day", "test"))

	var angka int = 89
	var booleanString string = "false"
	var floatString string = "3.14"
	var intString string = "8989"
	var boolValue bool = true

	fmt.Println(strconv.Itoa(angka))
	fmt.Println(strconv.ParseBool(booleanString))
	fmt.Println(strconv.ParseFloat(floatString, 10))
	fmt.Println(strconv.ParseInt(intString, 10, 64))
	fmt.Println(strconv.FormatBool(boolValue))
	fmt.Println(strconv.FormatInt(898, 10))

	currDate := time.Now()
	fmt.Println(currDate)
	fmt.Println("Year : ", currDate.Year())
	fmt.Println("Month : ", currDate.Month())
	fmt.Println("Day : ", currDate.Day())
	fmt.Println("Hour : ", currDate.Hour())
	fmt.Println("Minute : ", currDate.Minute())
	fmt.Println("Second : ", currDate.Second())

	var dateString string = "2023-08-15"
	parse, _ := time.Parse("2006-01-02", dateString)
	fmt.Println("Parse String to date: ", parse)

	name := "Audie"
	if length := len(name); length > 10{
		fmt.Println("Test fullfiled if conditional")
	}else{
		fmt.Println("Test else conditional")
	}

	switch name {
	case "Audie":
		fmt.Println("Hello Audie")
	case "Milson":
		fmt.Println("Hello Milson")
	default:
		fmt.Println("Hello ???")
	}	

	counter := 1
	fmt.Println("While Loop eq.")
	for counter <= 7{
		fmt.Println("Loop No. ", counter)
		counter++
	}
	fmt.Println("For Loop eq.")
	for i := 0; i < 10; i++ {
		fmt.Println("Loop No. ", i)
	}
	fmt.Println("Foreach eq.")
	sampleSlice := []string{"Ichigo", "Kurosaki"}
	for index, name := range sampleSlice{
		fmt.Println("Index: ", index, "=", name)
	}

	person := map[string]string{
		"name" : "Ichigo",
		"address" : "Soul Society",
	}
	
	fmt.Println(person)
	person["title"] = "Soul Reaper"
	fmt.Println(person, len(person))
	fmt.Println(person["name"])

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Audie Milson"
	book["wrong"] = "Ups"

	fmt.Println("Before", book)
	delete(book, "wrong")
	fmt.Println("After", book)

	namaDepan, namaBelakang := getCompleteName()
	fmt.Println(namaDepan, namaBelakang)

	kumpulanAngka := []int{10, 15, 20, 25}
	fmt.Println(sumAll(kumpulanAngka...))

	sayHelloWithFilter("test", spamFilter)
	sayHelloWithFilter("Ichigo", spamFilter)

	var newPerson map[string]string = nil
	fmt.Println(newPerson)

	newPerson = newMap("test")
	fmt.Println(newPerson)

	var cust Customer
	cust.Name = "Ichigo"
	cust.Address = "Rukongai"
	cust.Age = 23
	cust.greetCustomer()
	fmt.Println(cust)

	cust1 := Customer{
		Name: "Test Name",
		Address: "Test",
		Age: 32,
	}
	cust1.greetCustomer()

	res, err := division(5, 0)
	if(err != nil){
		fmt.Println("Error: ", err.Error())
	}else{
		fmt.Println(res)
	}

	runApplication(10)
	runAppPanic(false)

	testPerson := Person{Name: "Testing"}
	SayHello(testPerson)

	var result interface {} = random()

	var resAssertions string = result.(string)
	fmt.Println(resAssertions)

	switch value := result.(type){
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println(reflect.TypeOf(value), value)
	}
	users := []User{
		{
			Name: "testing 1",
			Age: 23,
		},
		{
			Name: "testing 2",
			Age: 17,
		},
		{
			Name: "testing 3",
			Age: 22,
		},
	}

	fmt.Println("Before: ", users)
	sort.Sort(UserSlice(users))
	fmt.Println("After: ", users)

	address1 := Address{
		City: "Subang",
		Province: "Jawa Barat",
		Country: "Indonesia",
	}

	address2 := &address1
	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(*address2)

	address3 := &Address{"Jakarta Barat", "DKI Jakarta", ""}
	ChangeAddressToIndonesia(address3)
	fmt.Println(*address3)
	
	var addr4 *Address = new(Address)
	fmt.Println(addr4)
	ChangeAddressToIndonesia(addr4)
	fmt.Println(*addr4)
	

	data := list.New()
	data.PushBack("Test")
	data.PushBack("Test1")
	data.PushBack("Test2")
	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	i := data.Front()

	fmt.Println("Loop Double Linkedlist value")
	for i != nil {
		fmt.Println(i.Value)
		i = i.Next()
	}

	fmt.Println("Loop Double Linkedlist value (in reverse)")
	for j := data.Back(); j != nil; j = j.Prev(){
		fmt.Println(j.Value)
	}

	fmt.Println("Circular Linkedlist")
	var data1 *ring.Ring = ring.New(5)
	for i := 0; i<data.Len(); i++{
		data1.Value = "Value" + strconv.FormatInt(int64(i), 10)
		data1 = data1.Next()
	}

	data1.Do(func(value interface{}){
		fmt.Println(value)
	})
}