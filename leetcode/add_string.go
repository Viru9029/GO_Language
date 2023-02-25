package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	a, err := strconv.Atoi(num1)
	checkNilError(err)
	b, err := strconv.Atoi(num2)
	checkNilError(err)
	return strconv.Itoa(a + b)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	num1 := "11"
	num2 := "123"
	fmt.Println(addStrings(num1, num2))
}
