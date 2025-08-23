package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var (
		a = 10
		b = 20
		c = a * b
	)

	fmt.Println(c)

	var i = 100
	i *= 20 // i = i * 20
	fmt.Println(i)

	i++ // i = i + 1
	fmt.Println(i)

	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)
}