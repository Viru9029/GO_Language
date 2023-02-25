package main

import (
	"fmt"
	"regexp"
	"strings"
)

func countSegments(s string) int {
	s = trim(s)
	if s == "" {
		return 0
	}
	str := strings.Split(s, " ")
	return len(str)
}

func trim(s string) string {
	result := strings.TrimSpace(s)
	regex := regexp.MustCompile("\\s+")
	result = regex.ReplaceAllString(result, " ")
	return result
}

func main() {
	senten := "Hello, my name is John"
	fmt.Println(countSegments(senten))
}
