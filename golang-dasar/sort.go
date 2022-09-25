package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func main() {
	users := []User{
		{"Audie", 21},
		{"Milson", 22},
		{"Ichigo", 18},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}

// Interface Sort
func (slice UserSlice) Len() int {
	return len(slice)
}

func (slice UserSlice) Less(i, j int) bool {
	return slice[i].Age < slice[j].Age
}

func (slice UserSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}