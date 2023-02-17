package main

import "fmt"

func reverseString(s []byte) []byte {
	i, j := 0, len(s)-1
	for i < j {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
		i++
		j--
	}
	return s
}

func main() {
	// str := "hello"
	// a := []byte{}
	// copy(a, str)
	//or
	str := "hello"
	a := make([]byte, len(str))
	copy(a, str)
	fmt.Println(reverseString(a))
}

// func reverseString(s []byte) []byte {
// 	res := []byte{}
// 	for i := len(s) - 1; i >= 0; i-- {
// 		res = append(res, s[i])
// 	}
// 	return res
// }
