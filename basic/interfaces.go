package main

import "fmt"

type i interface{}

func found(i interface{}) {
	fmt.Printf("Type = %T , value = %v\n", i, i)
}

func main() {
	s := "Welcome to Infinite CS"
	found(s)
	v := 1031217
	found(v)
}
