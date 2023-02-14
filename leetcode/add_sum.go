package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		var start = nums[i]
		for j := i + 1; j < len(nums); j++ {
			var end = nums[j]
			if (start + end) == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("Index nos : ", twoSum(nums, target))
}
