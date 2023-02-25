package main

import "fmt"

func Bubble_Sort(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
}

func main() {
	a := []int{1, 2, 2}
	fmt.Println(Bubble_Sort(a))
}
