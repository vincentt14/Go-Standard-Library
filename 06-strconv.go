package main

import (
	"fmt"
	"strconv"
)

// dokumentasi packaage strconv : https://pkg.go.dev/strconv

func main6(){
	 // string ke int = Atoi
	 // int ke string = Itoa

	// b, err := strconv.ParseBool("true")
	// f, err := strconv.ParseFloat("3.1415", 64)
	// i, err := strconv.ParseInt("-42", 10, 64)
	// u, err := strconv.ParseUint("42", 10, 64)
	
	result, err := strconv.ParseBool("true")
	if err != nil{
		fmt.Println("Error", err.Error())
	}else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("1000")
	if err != nil{
		fmt.Println("Error", err.Error())
	}else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var startInt string = strconv.Itoa(999)
	fmt.Println(startInt)
}