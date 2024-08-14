package main

import "fmt"

//pointer - type that stores memory location

func main() {
	// var p *int32 = new(int32)
	// *p = 10
	// // when we initalize the pointer with new(), get to the free memory location and initalized it to zero 
	// // anssinged pointer has value as nil
	// var i int32
	// fmt.Printf("The value p points to is: %v", *p)
	// fmt.Printf("\nThe value if i is: %v\n", i)

	// p = &i
	// *p = 1
	// fmt.Printf("\nThe value of i is: %v\n", i)

	// slice uses pointer under the hood


	var thing1 = [5]float64{1,2, 3,4,5}
	fmt.Printf("The memory location of the thing1 array is: %p\n", &thing1)
	// var result [5]float64 = square(thing1)

	var result [5]float64 = pointerSquare(&thing1)
	fmt.Printf("The result is: %v \n", result)
	fmt.Printf("The result is: %v \n", thing1)

}

// func square(thing2 [5]float64) [5]float64{
// 	fmt.Printf("The memory location of the thing2 array is: %p\n", &thing2)
// 	for i := range thing2{
// 		thing2[i] = thing2[i]*thing2[i]
// 	}
// 	return thing2
// }

func pointerSquare(thing2 *[5]float64) [5]float64{
	fmt.Printf("The memory location of the thing2 array is: %p\n", &thing2)
	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return *thing2
}