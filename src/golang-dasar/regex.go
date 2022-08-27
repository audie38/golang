package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`[A-Za-z]`)

	fmt.Println(regex.MatchString("Audie"))
	fmt.Println(regex.MatchString("ichigo"))
	fmt.Println(regex.MatchString("88888"))
	fmt.Println(regex.FindAllString("Audie Milson", 20))
	fmt.Println(regex.FindAllString("Audie Milson", -1))
}