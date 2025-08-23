package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Ariyan"
	names[1] = "Andryan"
	names[2] = "Aryja"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	var values = [3]int{
		90, 80, 70,
	}
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}