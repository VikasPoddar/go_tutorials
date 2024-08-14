package main

import (
	"fmt"
	"strings"
)
func main() {
	// sting, runes and bytes
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v %T\n", indexed, indexed)
	// ASCII uses 7bit encoding for repesentation 
	// UTF-32 uses 32bit encoding for reperesentation 
	// UTF-8 uses variable length encoding, uses code-point for adding larger representation
	// 0xxxxxxx (1byte) code point first: 0 last: 127 (regular ASCII)
	// 110xxxxx | 10xxxxxx (2byte) code point first: 128 last: 2047
	// é - 233 (11101001) -> (PADDING 00011101001 it) -> 110xxxxx | 10xxxxxx) 
	// golang uses utf-8 encoding for representing string 
	// for index, value := range myString{
	// 	fmt.Println(index, value)
	// }
	// taking lenght of an array is useless in golang 
	fmt.Printf("The length of 'myString' is %v\n", len(myString))
	// hence for using string for iteration, we need to cast string into datatype runes
	var iteratableString = []rune(myString)
	for index, value := range iteratableString{
		fmt.Printf("%v %v \n", index, value)
	}

	//string building
	var strSlice = []string{"v", "i", "k", "a", "s"}
	// var catStr = ""
	// for index := range strSlice{
	// 	catStr += strSlice[index]
	// }
	// fmt.Printf("%v\n", catStr)
	// string in go are immutable 
	var strBulider strings.Builder
	for index := range strSlice{
		strBulider.WriteString(strSlice[index])
	}
	var catStr = strBulider.String()
	fmt.Printf("%v\n", catStr)
	
}