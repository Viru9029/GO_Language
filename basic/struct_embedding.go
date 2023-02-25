package main

import "fmt"

type describer interface {
	describe() string
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func printss(m describer) {
	fmt.Println(m.describe())
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("%v, ,%v\n", co.num, co.str)
	fmt.Println(co.base.num)
	fmt.Println("describe : ", co.describe())
	bo := base{num: 322}
	printss(bo)

	var d describer = co
	fmt.Println("describer : ", d.describe())
}
