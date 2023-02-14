package main

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString("\n")
	// fmt.Println(reader)
	var num int
	fmt.Printf("Enter the number : ")
	fmt.Scanf("%d", &num)
	fmt.Printf("Entered no is : %d\n", num)
	fmt.Println(split(num))
}
