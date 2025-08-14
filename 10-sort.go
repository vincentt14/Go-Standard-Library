package main

import (
	"fmt"
	"sort"
)

// dokumentasi packaage sort : https://pkg.go.dev/sort

type User struct {
	Name string
	Age  int
}

type UserSclie []User

func (s UserSclie) Len() int {
	return len(s)
}

func (s UserSclie) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSclie) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main10() {
	users := []User{
		{"Eko", 30},
		{"Budi", 35},
		{"Joko", 25},
		{"Adit", 20},
	}

	sort.Sort(UserSclie(users))
	fmt.Println(users)
}
