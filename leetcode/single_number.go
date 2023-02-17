package main

import "fmt"

func singleNumber(nums []int) int {
	count_store := make(map[int]int)
	for _, key := range nums {
		count_store[key]++
	}
	res := 0
	for _, key := range nums {
		if count_store[key] == 1 {
			res += key
		}
	}
	return res
}

func main() {
	array := []int{2, 2, 2, 1}
	fmt.Println(singleNumber(array))
}
