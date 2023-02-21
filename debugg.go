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
