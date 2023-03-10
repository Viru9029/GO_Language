package main

import (
	"fmt"
	"strings"
)

func mostWordsFound(sentences []string) int {
	count := []int{}
	for i := 0; i < len(sentences); i++ {
		str_Split := len(strings.Split(sentences[i], " "))
		count = append(count, str_Split)
	}
	for i := 0; i < len(count); i++ {
		for j := i + 1; j < len(count); j++ {
			if count[i] < count[j] {
				count[j], count[i] = count[i], count[j]
			}
		}
	}
	return count[0]
}

// func mostWordsFound(sentences []string) int {
//     Max_length:= len(strings.Split(sentences[0]," "))
//     for _,val := range sentences {
//         length:=len(strings.Split(val," "))
//         if length > Max_length {
//             Max_length = length
//         }
//     }
//     return Max_length
// }

func main() {
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(sentences))
}
