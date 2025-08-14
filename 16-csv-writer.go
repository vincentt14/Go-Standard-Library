package main

import (
	"encoding/csv"
	"os"
)

func main16() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"eko", "Kurniawan", "khannedy"})
	_ = writer.Write([]string{"vudi", "pratama", "nugraha"})
	_ = writer.Write([]string{"joko", "marro", "diah"})
	
	writer.Flush()
}