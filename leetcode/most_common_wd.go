package main

import (
	"fmt"
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	lower_case := strings.ToLower(paragraph)
	lower_case = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(lower_case, "")
	split_wd := strings.Split(lower_case, " ")
	count_wd := make(map[string]int)
	for i := 0; i < len(split_wd)-1; i++ {
		count_wd[split_wd[i]] += 1
	}
	if 
	return count_wd
}

func main() {
	str := "Bob hit a ball, the hit BALL flew far after it was hit."
	bann := []string{"hit"}
	// split_wd := strings.Split(str, " ")
	fmt.Println(mostCommonWord(str, bann))
}
