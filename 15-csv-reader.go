package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main15() {
	csvString := "eko kurniawan khannedy\n" + 
							"budi pratama nugraha\n" + 
							"joko morror diah"

	reader := csv.NewReader(strings.NewReader(csvString))
	
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

	fmt.Println(record)
	}
}