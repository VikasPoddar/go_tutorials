package main

// array,slices and maps
// looping control structure 

// Arrays - Fixed length, Same Type, Indexable, Contiguous in Memory

import "fmt"

func main() {
	// var intArr [3]int32
	// intArr[1] = 123
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:3])
	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	// var intArr [3]int32 = [3]int32{1,2,3}
	// fmt.Println(intArr)
	intArr := [3]int32{1, 2, 3}
	// fmt.Println(intArr)


	// slices are warp arrays - more general, powerful, convenient interface to seq. of data

	// var intSlice []int32 = []int32{4, 5, 6}
	// fmt.Printf("The length of %v with capacity %v", len(intSlice), cap(intSlice))
	// intSlice = append(intSlice, 7)
	// fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))
	// if we try to access values outside the assigned valuewe would get a index out of bound error 
	// var intSlice2 []int32 = []int32{8, 9}
	// intSlice = append(intSlice, intSlice2...) // spread operator
	// fmt.Println(intSlice)

	// var intSlice3 []int32 = make([]int32, 3, 8) // length and capacity, by default the length == capacity
	// fmt.Println(intSlice3)

	// maps - key value pair
	// var myMap map[string]uint8 = make(map[string]uint8)
	// fmt.Println(myMap)

	myMap2 := map[string]uint8{"Adam": 8, "Sarah": 45}
	// fmt.Println(myMap2)
	// fmt.Println(myMap2["Adam"])

	var age, ok = myMap2["Adam"]
	if ok{
		fmt.Printf("The age is %v\n", age)
	}else{
		fmt.Println("Invalid Name")
	}

	// iteration 
	for name := range myMap2{
		fmt.Printf("Name: %v\n", name)
	}

	for name, age := range myMap2{
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	for index, value := range intArr{
		fmt.Printf("Index: %v, Value: %v \n", index, value)
	}

	// using for loop as a while loop in golang 
	// as while loop does not exists in golang

	// var index int = 0
	// for index <= 10{
	// 	fmt.Println(index)
	// 	index +=1
	// }
// 	for{
// 		if index >10{
// 			break
// 		}
// 		fmt.Println(index)
// 		index += 1
// 	}
	for index := 0; index <= 10; index++ {
		fmt.Println(index)
	}
}