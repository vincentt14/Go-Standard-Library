package main

import (
	"fmt"
	"os"
)

// dokumentasi packaage os : https://pkg.go.dev/os

func main3(){
	args := os.Args
	fmt.Println(args)

	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil{
		fmt.Println(hostname)
	}else {
		fmt.Println("error", err)
	}
}