package main

import "fmt"

func main() {
	var arr1 [5]int32 = [5]int32{1, 2, 3, 4, 5}
	fmt.Printf("Memory location of arr1 is :%p", &arr1)
	var result [5]int32 = square(arr1)
	fmt.Printf("Result is: %v", result)
	//In this example both the arr point to different memory locations increasing memory usage
	// we can instead use pointers and modify the original array rather than making copies of it
	var a [5]int32 = [5]int32{3, 4, 5, 6, 7}
	fmt.Printf("Memory location of a is :%p", &a)
	var res [5]int32 = squareWithPointers(&a)
	fmt.Printf("Result is %v", res)
}

func square(arr2 [5]int32) [5]int32 {
	fmt.Printf("Memory location of arr 2 is : %p", &arr2)
	for i := range arr2 {
		arr2[i] = arr2[i] * arr2[i]
	}

	return arr2
}

func squareWithPointers(arr *[5]int32) [5]int32 {
	fmt.Printf("Memory location of arr is : %p", arr)
	for i := range arr {
		arr[i] = arr[i] * arr[i]
	}
	return *arr
}
