package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct{
	Street, Country, PostalCode string
}

type Customer struct{
	FirstName string
	MiddleName string
	LastName string
	Age int
	IsHeadCaptain bool
	Hobbies []string
	Addreses []Address
}

type Product struct{
	Id string	`json:"id"`
	Name string	`json:"name"`
	ImageURL string	`json:"image_url"`
}

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T){
	logJson("Genryusai")
	logJson(1)
	logJson(true)
	logJson([]string{"Shigekuni", "Genryusai", "Yamamoto"})
	customer := Customer{
		FirstName: "Yamamoto",
		MiddleName: "Shigekuni",
		LastName: "Genryusai",
		Age: 1000,
		IsHeadCaptain: true,
		Hobbies: []string{"Lead", "Train", "Manage"},
		Addreses: []Address{
			{
				Street: "1st Division Barrack",
				Country: "Soul Society",
				PostalCode: "00001",
			},
			{
				Street: "Rukongai District 38",
				Country: "Soul Society",
				PostalCode: "00038",
			},
		},
	}
	logJson(customer)
}

func TestDecode(t *testing.T){
	jsonRequest := `{"FirstName":"Yamamoto","MiddleName":"Shigekuni","LastName":"Genryusai","Age":1000,"IsHeadCaptain":true,"Hobbies":["Lead","Train","Manage"],"Addreses":[{"Street":"1st Division Barrack","Country":"Soul Society","PostalCode":"00001"},{"Street":"Rukongai District 38","Country":"Soul Society","PostalCode":"00038"}]}`
	jsonBytes := []byte(jsonRequest)
	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(customer)
}

func TestEncodeTag(t *testing.T){
	product := Product{
		Id: "P001",
		Name: "Product One",
		ImageURL: "Product One Image",
	}

	logJson(product)
}

func TestDecodeTag(t *testing.T){
	jsonReq := `{"id":"P001","name":"Product One","image_url":"Product One Image"}`
	jsonBytes := []byte(jsonReq)
	product := &Product{}
	err := json.Unmarshal(jsonBytes, product)
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(product)
}