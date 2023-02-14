package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of pizza : ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for string.", input)
	fmt.Printf("Type of this rating is %T\n ", input)
}
