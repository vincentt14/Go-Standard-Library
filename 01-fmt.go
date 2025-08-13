package main

import "fmt"

// dokumentasi packaage fmt % : https://pkg.go.dev/fmt#pkg-index

func main1() {
	firstName := "Eko"
	lastName := "Khannedy"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}