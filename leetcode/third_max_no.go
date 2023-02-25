package main

import "fmt"

func thirdMax(nums []int) int {
	slice_Input := nums
	for i, val := range slice_Input {
		for k := i + 1; k < (len(nums) - 1); k++ {
			fmt.Println(val, nums[k])
			if val == nums[k] {
				slice_Input = append(slice_Input[:k], slice_Input[k+1:]...)
			}
		}
	}
	fmt.Println(slice_Input)
	for i := 0; i < len(slice_Input); i++ {
		for j := i + 1; j < len(slice_Input); j++ {
			if slice_Input[j] > slice_Input[i] {
				slice_Input[j], slice_Input[i] = slice_Input[i], slice_Input[j]
			}
		}
	}
	fmt.Println(slice_Input)
	if len(slice_Input) >= 3 {
		return slice_Input[2]
	} else if len(slice_Input) <= 2 {
		return slice_Input[0]
	}
	return 0
}

func main() {
	a := []int{1, 2, 2}
	fmt.Println(thirdMax(a))
}
