package main

import (
	"fmt"
	"errors"
)

func main() {
	// Below is the list of data types in go
	// int8, int16, int32, int64, uint32
	// float32, float64
	// string
	// boolean
	// error

	var num int8 = 100
	var num2 int16 = 10000
	var num3 int32 = 1000000000
	num4 := 1000000000000000000

	fmt.Println(num, num2, num3, num4)


	var numfloat float32 = 123.45
	var numfloat2 float64 = 1230000000.45
	numfloat3 := 12345.67

	fmt.Println(numfloat, numfloat2, numfloat3)


	var str string = "text"
	str2 := "text2"

	fmt.Println(str, str2)


	var booleano bool = true
	var booleano2 bool // value will be false if no value is being passed
	
	fmt.Println(booleano, booleano2)


	var erro error = errors.New("This is an error")
	fmt.Println(erro)
}