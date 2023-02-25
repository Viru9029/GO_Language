package main

import "fmt"

func majorityElement(nums []int) int {
	storeCount := make(map[int]int)
	greater := 0
	res := 0
	for _, val := range nums {
		storeCount[val]++
		if storeCount[val] > greater {
			greater = storeCount[val]
			res = val
		}
	}
	return res

}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
