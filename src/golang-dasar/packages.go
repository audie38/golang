package main

import (
	"flag"
	"fmt"
	"os"
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

}