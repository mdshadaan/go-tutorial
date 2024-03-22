package main

import (
	"errors"
	"fmt"
)

func main() {
	var result, err = multiply(3, 4)
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case result >= 1000:
		fmt.Printf("Big number %v", result)
	default:
		fmt.Printf("small number!")
	}
}

func multiply(num1 int, num2 int) (int, error) {
	var err error
	if num1 == 0 || num2 == 0 {
		err = errors.New("Make sure both numbers are not zero")
		return 0, err
	}
	return num1 * num2, err
}
