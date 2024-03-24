package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr[0])
	//slices are a wrapper around arrays and gives us more functionality
	var intSlice []int32 = []int32{4, 5, 6}
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)

	//we can also append multiple values to a slice by using the spread operator
	var intSlice2 []int32 = []int32{8, 9}
	intSlice2 = append(intSlice2, intSlice...)
	fmt.Println(intSlice2)

	var myMap map[string]uint8 = make(map[string]uint8)
	//In Go , map will always return a value even if the key does not exist.
	// it also returns another varible indicating whether a key exists or not
	fmt.Println(myMap)
	var myMap2 map[string]uint8 = map[string]uint8{"shadaan": 27, "nic": 25}
	var age, ok = myMap2["zoheb"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("Key does not exist")
	}
	//In Go , for is the only looping construct available.
	for key, val := range myMap2 {
		fmt.Printf("Name: %v , age: %v", key, val)
	}

	for index, val := range intSlice2 {
		fmt.Printf("Index: %v , value: %v", index, val)
	}

}
