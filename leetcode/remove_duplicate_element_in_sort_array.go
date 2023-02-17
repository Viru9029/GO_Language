package main

import "fmt"

func removeDuplicates(nums []int) int {
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[res] != nums[i] {
			res++
			nums[res] = nums[i]
		}
	}
	return res + 1
}

func main() {
	nums := []int{-3, -1, 0, 0, 0, 3, 3}
	fmt.Println(removeDuplicates(nums))
}

//output 3

//my logic
//func removeDuplicates(nums []int) int {
// 	key_val_store := make(map[int]int)
// 	res := 0
// 	for _, val := range nums {
// 		key_val_store[val] += 1
// 	}
// 	for key := range key_val_store {
// 		nums[res] = key
// 		res += 1
// 	}
// 	fmt.Println(nums)
// 	return res
// }
