package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"` //Struct Tag
}

func main() {
	sample := Sample{"Audie"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0).Name

	fmt.Println(sampleType.NumField())
	fmt.Println(structField)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	fmt.Println(sampleType.Field(0).Tag.Get("min"))
	fmt.Println(IsValid(sample))

}

func IsValid(data interface{}) bool{
	t := reflect.TypeOf(data)
	for i := 0; i< t.NumField(); i++{
		field := t.Field(i)
		if field.Tag.Get("required") == "true"{
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}

	return true
}