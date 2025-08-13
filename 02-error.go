package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

// dokumentasi package errors : https://pkg.go.dev/errors

func main2() {
	err := getById("")
	if err != nil {
		if errors.Is(err, ValidationError){ // errors.Is untuk caritau jenis error
			fmt.Println("validation error")
		}else if errors.Is(err, NotFoundError){
			fmt.Println("not found error")
		}else {
			fmt.Println("unknown error")
		}
	}
}

func getById(id string) error {
	if id == ""{
		return ValidationError
	}

	if id != "Eko"{
		return NotFoundError
	}

	return nil
}