package main

import (
	"errors"
	"fmt"
)

func main(){
	var myName string = "Vikas Poddar"
	const NAME string = "Harshit Poddar"
	printMe()
	printName("No Name")
	printName(myName)
	printName(NAME)
	var numerator int = 10
	var denominator int = 0
	var answer, remainder, err = intDivision(numerator, denominator)
	// if err!=nil{
	// 	fmt.Println(err.Error())
	// }else if remainder == 0{
	// 	fmt.Printf("The result of the integer division is %v", answer)
	// }else{
	// 	fmt.Printf("The result of the division operation on %v by %v are as follows: \n quotient: %v remainder: %v", numerator, denominator, answer, remainder )
	// }
	switch{
	case err!=nil:
		fmt.Println(err.Error())
	case remainder==0:
		fmt.Printf("The result of the integer division is %v", answer)
	default:
		fmt.Printf("The result of the division operation on %v by %v are as follows: \n quotient: %v remainder: %v", numerator, denominator, answer, remainder)
	}
}

func printMe(){
	fmt.Println("Printing value")
}

func printName(name string){
	fmt.Println("Hello, ",name)
}

func intDivision(numerator int, denominator int) (int, int, error){
	var err error
	if denominator==0{
		err = errors.New("cannot divide by Zero")
		return 0, 0, err 
	}
	var result int = numerator/denominator
	var reminder int = numerator%denominator
	return result, reminder, err
}