package main

import "fmt"

func countVowelSubstrings(word string) int {
	count := 0
	//word := []byte(word)
	for i := 0; i < len(word)-1; i++ {
		if word[i] == "a" || word[i] == 'e' || word[i] == "i" || word[i] == "o" || word[i] == "u" {
			count += 1
		}
		return count
	}
}

func main() {
	str := "aeiouu"
	fmt.Println(countVowelSubstrings(str))
}
