package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// dokumentasi packaage ring : https://pkg.go.dev/ring

func main9(){
	data := ring.New(5)

	// manual
	// data.Value = "value 1"

	// data = data.Next()
	// data.Value = "value 2"
	
	// data = data.Next()
	// data.Value = "value 3"
	
	// data = data.Next()
	// data.Value = "value 4"
	
	// data = data.Next()
	// data.Value = "value 5"

	// pakai for (otomatis)
	for i := 0; i < data.Len(); i++{
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(value any){
		fmt.Println(value)
	})
}