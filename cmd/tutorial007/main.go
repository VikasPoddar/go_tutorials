package main

// Structs and Interfaces

import (
	"fmt"
)

// type gasEngine struct{
// 	mpg uint8
// 	gallons uint8
// 	ownerInfo owner
// }

type gasEngine struct{
	mpg uint8
	gallons uint8
	owner
	//int we can have nameless field
}

type owner struct{
	name string
}

// method
func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

type electricEngine struct{ 
	mpkpw uint8
	khw uint8
	owner
}

func (e electricEngine) milesLeft() uint8{
	return e.mpkpw*e.khw
}

type engine interface{
	milesLeft() uint8
}


func canMakeIt(e engine, miles uint8){
	if miles <= e.milesLeft(){
		fmt.Println("You can make it there!")
	}else{
		fmt.Println("Need to fuel up first!")
	}
}
// default values of each field zero

func main(){
	// var myEngine gasEngine = gasEngine{mpg: 25, gallons: 45} // struct literal syntax
	// var myEngine gasEngine = gasEngine{25, 15, owner{"Vikas"}}
	var myEngine = gasEngine{25, 10, owner{"Vikas"}}
	var myEngine2 = electricEngine{25, 100, owner{"Harshit"}}
	canMakeIt(myEngine, 100)
	canMakeIt(myEngine2, 100)
	// myEngine.gallons = 100
	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
	//anonymous struct declear and initalize same time
	// non-reusable
	// var myEngine2 = struct{
	// 	mpg uint8
	// 	gallons uint8
	// }{25, 15}
	// fmt.Println(myEngine2)
	// fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())
}