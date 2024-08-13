package main

import (
	"fmt"
	"unicode/utf8"
)
func main() {
	var intNum int
	var floatNum float32
	// int type
	// int8
	// int16
	// int32
	// int64
	// uint 
	// uint8
	// uint16
	// uint16
	// uint32
	// uint64
	// float type
	// float32
	// float64
	// arithmetic operations
	// + - / *
	// boolean
	var myTrueBoolean bool = true
	var myFalseBoolean bool = false
	fmt.Println(myTrueBoolean, " ", myFalseBoolean)
	// bitwise-boolean operations
	// relational operations
	
	// no mixed type operations
	// floatNum = 3.13 + intNum
	floatNum = 3.13 + float32(intNum)
	// explicit type casting before mixed operations
	// integer division returns rounded-down integer value
	fmt.Print(intNum, floatNum, "\n")
	var myLongString string = "Hello, I'm Vikas Poddar and" + " " + "I'm learning go lang in my freetime"
	fmt.Println(myLongString)
	var normalString string = "A"
	var unicodeString string = "γ"
	
	// but the len function returns the length in bytes
	fmt.Println("Normal String: ", normalString,"length using built-in len function: ", len(normalString), "\nUnicode String: ", unicodeString, "Length using built-in len function: ", len(unicodeString))
	fmt.Println("Normal String: ", normalString,"length using RuneCountInString function: ", utf8.RuneCountInString(normalString), "\nUnicode String: ", unicodeString, "length using RuneCountInString function: ", utf8.RuneCountInString(unicodeString))
	
	// NOTE: Rune is a built-in struct for holding in uni-codes
	var myRune rune = 'a'
	fmt.Println(myRune)
	// declaration & initialization 
	// declaration without initialization get assigned to a default value
	// uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, rune - 0
	// string - ''
	// boolean - false
	// style of declaration & initialization
	// style 1:
	// var myInt int = 10
	// style 2:
	// var myInt = 10
	// style 3:
	// var myInt int
	// myInt = 10
	// style 4: 
	// myInt := 10
	// multiple initiation
	// var myInt, my-float int = 10, 3
	// fmt.Println(myInt, myfloat)
	// constant 
	const myConstValue string = "Constant Value"
	fmt.Println(myConstValue)
	} 