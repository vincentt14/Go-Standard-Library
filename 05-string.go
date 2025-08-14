package main

import (
	"fmt"
	"strings"
)

// dokumentasi packaage strings : https://pkg.go.dev/strings

func main5(){
	 fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	 fmt.Println(strings.Split("Eko Kurniawan", " "))
	 fmt.Println(strings.ToLower("Eko Kurniawan"))
	 fmt.Println(strings.ToUpper("Eko Kurniawan"))
	 fmt.Println(strings.Trim("      Eko Kurniawan    ", " "))
	 fmt.Println(strings.ReplaceAll("Eko Kurniawan", "Eko", "Budi"))
	}