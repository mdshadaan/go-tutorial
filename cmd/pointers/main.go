package main

import "fmt"

func main() {
	//initialising a pointer
	var p *int32 = new(int32)
	var i int32
	// the * notation serves two functions - it can be used to dereference a pointer like *p will get the value stored at p
	//we can also use it to assign value to a pointer like *p = 25
	fmt.Printf("The value p points to is : %v", *p)
	fmt.Printf("\n The value of i is %v", i)

	//changing value stored at variable p
	*p = 25
	fmt.Printf("\n Value changed of var p : %v", *p)

	i = 35
	p = &i //p now refers to memory address of i
	fmt.Printf("\n Value of p : %v", *p)
}
