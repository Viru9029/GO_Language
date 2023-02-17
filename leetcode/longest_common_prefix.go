package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	for index, val := range strs[0] {
		for _, j := range strs[1:] {
			if index == len(j) || byte(val) != j[index] {
				return strs[0][:index]
			}
		}
	}
	return strs[0]
}

func main() {
	a := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(a))
}

//func longestCommonPrefix(strs []string) string {
// 	res := ""
// 	leng := []int{}
// 	leng = append(leng, len(strs[0]))
// 	leng = append(leng, len(strs[1]))
// 	leng = append(leng, len(strs[2]))
// 	sort.Ints(leng)
// 	fmt.Println(leng)
// 	for i := 0; i < leng[0]-1; i++ {
// 		if string(strs[0][i]) == string(strs[1][i]) && string(strs[1][i]) == string(strs[2][i]) {
// 			res += string(strs[0][i])
// 		} else {
// 			return res
// 		}
// 	}
// 	return ""
// }
