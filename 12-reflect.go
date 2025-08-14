package main

import (
	"fmt"
	"reflect"
)

// struct tag
type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("type name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)
		
		fmt.Println(structField.Name, "with type", structField.Type)

		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

// dokumentasi packaage reflect : https://pkg.go.dev/reflect

func main12() {
	readField(Sample{"Eko"})
	readField(Person{"Eko", "jajaja", "dada@gmail.com"})


	person := Person{
		Name: "Eko",
		Address : "Jakarta",
		Email: "",
	}

	fmt.Println(isValid(person))
}

func isValid(value any)(result bool){

	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++{
		f := t.Field(i)
		if f.Tag.Get("required") == "true"{
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}

	return result
}