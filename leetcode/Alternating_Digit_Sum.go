package main

import "fmt"

func alternateDigitSum(n int) int {
	sep_Numbers := []int{}
	for n > 0 {
		num := n % 10
		sep_Numbers = append(sep_Numbers, num)
		n = n / 10
	}

	reverse_Number := []int{}
	for i := len(sep_Numbers) - 1; i >= 0; i-- {
		reverse_Number = append(reverse_Number, sep_Numbers[i])
	}

	res := 0
	for k := 0; k < len(reverse_Number); k++ {
		if k%2 == 0 {
			res = res + reverse_Number[k]
		} else {
			res = res - reverse_Number[k]
		}
	}
	return res
}

func main() {
	n := 521
	fmt.Println(alternateDigitSum(n))
}
