package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	result := strings.Split(s, " ")
	return len(result[len(result)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("         Hello World            "))
}
