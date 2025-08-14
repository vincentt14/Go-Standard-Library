package main

import (
	"fmt"
	"slices"
)

func main17() {
	names := []string{"Joki", "Eko", "Cent", "Ringo"}
	values := []int{100, 80, 30, 70}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Eko"))
	fmt.Println(slices.Index(names, "Cent"))
	fmt.Println(slices.Index(names, "dsadas"))
}