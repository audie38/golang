package main

import (
	"flag"
	"fmt"
	"os"
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
}