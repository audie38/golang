package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main(){

	// Package OS 
	fmt.Println(os.Args)

	name, err  := os.Hostname()
	if err == nil{
		fmt.Println(name)
	}else{
		fmt.Println(err.Error())
	}

	// Package Flag
	// go run src/golang-dasar/packages.go -host=localhost -user=root -password=root
	var host *string = flag.String("host", "localhost", "Put your db host")
	flag.Parse()
	fmt.Println(host)
	fmt.Println(*host)

	// Package String
	fmt.Println(strings.Contains("Audie Milson", "Audie"))
	fmt.Println(strings.Split("Audie Milson", " "))
	fmt.Println(strings.ToLower("Audie Milson"))
	fmt.Println(strings.ToUpper("Audie Milson"))
	fmt.Println(strings.Trim("               Audie  Milson            ", " "))
	fmt.Println(strings.ReplaceAll("Audie Audie Audie", "Audie", "Milson"))

	// Package strconv
	fmt.Println("Package strconv")
	boolean, err := strconv.ParseBool("true")
	if err == nil{
		fmt.Println(boolean)
	}else{
		fmt.Println("Error: ", err.Error())
	}

	number, err := strconv.ParseInt("80", 10, 64) // convert from string : input, base, return bit size
	if err == nil{
		fmt.Println(number)
	}else{
		fmt.Println("Error: ", err.Error())
	}

	value := strconv.FormatInt(1000000, 10) // convert to string : input, base
	fmt.Println(value)

	valInt, err := strconv.Atoi("2000000") // konversi string ke number
	fmt.Println(valInt)
	valString := strconv.Itoa(8000) // konversi ke string
	fmt.Println(valString)

	// Package math
	fmt.Println("Package Math")
	fmt.Println(math.Round(8.6))
	fmt.Println(math.Round(8.3))
	fmt.Println(math.Floor(8.6))
	fmt.Println(math.Ceil(8.3))
	fmt.Println(math.Max(8, 12))
	fmt.Println(math.Min(8, 12))

}