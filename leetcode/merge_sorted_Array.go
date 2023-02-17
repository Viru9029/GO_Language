package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--

	for i := m + n + 1; i >= 0 && n >= 0; i-- {
		if m < 0 || nums1[m] < nums2[n] {
			nums1[i] = nums2[n]
			n--
		} else {
			nums1[i] = nums1[m]
			m--
		}
	}
}

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := 3
	c := []int{2, 5, 6}
	d := 3
	fmt.Println(merge(a, b, c, d))
}
