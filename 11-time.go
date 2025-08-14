package main

import (
	"fmt"
	"time"
)

// dokumentasi packaage time : https://pkg.go.dev/time

func main11(){
	now := time.Now()
	fmt.Println(now) // timezone komputer
	fmt.Println(now.Local()) // timezone kita

	utc := time.Date(2009, time.November, 10,20,0,0,0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"

	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("err", err.Error())
	}else {
		fmt.Println(valueTime)
		fmt.Println(valueTime.Day())
		fmt.Println(valueTime.Month())
		fmt.Println(valueTime.Year())
	}

	// duration
	duration1 := 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond 
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration3)
	fmt.Println("duration : %d", duration3)
}