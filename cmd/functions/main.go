package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var sum int = add(4, 5)
	fmt.Println(sum)
	var length, upperCase, err = lenAndUpperCase("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Length of string is %v and uppercase : %v ", length, upperCase)
	}
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

// we can have functions that return more than one value
// Convention in go is to also return error (which is a datatype also)
// when we know our function can cause certain issues.
func lenAndUpperCase(str string) (int, string, error) {
	var length int = utf8.RuneCountInString(str)
	var err error
	if length == 0 {
		err = errors.New("Pass in a valid string!")
		return 0, "", err
	}
	upperCase := strings.ToUpper(str)
	return length, upperCase, err
}
