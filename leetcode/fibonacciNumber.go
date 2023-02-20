package main

import "fmt"

func fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		return (fib(n-1) + fib(n-2))
	}

}

func main() {
	n := 25
	fmt.Println(fib(n))
}
