package main

import (
	"container/list"
	"fmt"
)

// dokumentasi packaage container/list : https://pkg.go.dev/container/list

func main8(){
	data := list.New()
	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")

	pertama := data.Front()
	kedua := pertama.Next()
	ketiga := pertama.Next().Next()

	fmt.Println(pertama.Value)
	fmt.Println(kedua.Value)
	fmt.Println(ketiga.Value)

	for o := data.Front(); o != nil; o = o.Next(){
		fmt.Println(o.Value)
	}
}