package main

import "fmt"

func mostWordsFound(sentences []string) int {
	count := []int{}
	for i := 0; i < len(sentences); i++ {
		count = append(count, len(sentences[i]))
	}
	fmt.Println(count)
	for i := 0; i < len(count); i++ {
		for j := i + 1; j < len(count); j++ {
			if count[i] < count[j] {
				count[j], count[i] = count[i], count[j]
			}
		}
	}
	return count[0]
}

func main() {
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(sentences))
}
