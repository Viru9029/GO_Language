package main

import "fmt"

// func removeElement(nums []int, val int) int {
// 	snowBallSize := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == val {
// 			snowBallSize++
// 		} else if snowBallSize > 0 {
// 			nums[i], nums[i-snowBallSize] = nums[i-snowBallSize], nums[i]
// 		}
// 	}

// 	return len(nums) - snowBallSize
// }

//Second correct code
func removeElement(nums []int, val int) int {
	replace_index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[replace_index] = nums[i]
			replace_index += 1
		}
	}
	return replace_index
}

//[3,2,2,3]
//[2,]

func main() {
	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
	b := 2
	fmt.Println(removeElement(a, b))
}

// func removeElement(nums []int, val int) int {
//     numbers := nums
//     for i:=0;i < len(nums);i++{
//         if nums[i] == val{
//             index := int(i)
//             numbers = append(numbers[:index+copy(numbers[index:])],numbers[index+1:]...)
//         }
//     }
//     return len(numbers)
// }

//correct code
// func removeElement(nums []int, val int) int {
// 	count := []int{}
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != val {
// 			count = append(count, nums[i])
// 		}
// 	}
// 	return len(count)
// }
