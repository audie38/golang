package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct{
	FirstName string
	MiddleName string
	LastName string
	Age int
	IsHeadCaptain bool
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
	}
	logJson(customer)
}

func TestDecode(t *testing.T){
	jsonRequest := `{"FirstName": "Ichigo", "MiddleName": "Kurosaki", "LastName" : "Shiba", "Age": 20, "IsHeadCaptain" : false}`
	jsonBytes := []byte(jsonRequest)
	customer := &Customer{}
	json.Unmarshal(jsonBytes, customer)
	fmt.Println(customer)
}