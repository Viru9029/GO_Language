package main

import "fmt"

func addToArrayForm(num []int, k int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	sum += k
	return sum

}

func main() {
	num := []int{1, 2, 0, 0}
	k := 34
	fmt.Println(addToArrayForm(num, k))
}
