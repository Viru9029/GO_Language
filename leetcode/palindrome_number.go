package main

import "fmt"

func isPalindrome(x int) bool {
	value := 0
	num := x
	for x > 0 {
		remainder := x % 10
		value *= 10
		value += remainder
		x /= 10
	}
	if value == num {
		return true
	}
	return false
}

func main() {
	fmt.Println("Number is palindrome : ", isPalindrome(121))
}

// Remainder :  1
// value :  0
// value2 :  1
// x :  121
// x 2 :  12
// Remainder :  2
// value :  10
// value2 :  12
// x :  12
// x 2 :  1
// Remainder :  1
// value :  120
// value2 :  121
// x :  1
// x 2 :  0
// Number is palindrome :  true
