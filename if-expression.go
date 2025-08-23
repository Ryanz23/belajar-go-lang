package main

import "fmt"

func main() {
	name := "Ariyan"

	if name == "Andryan" {
		fmt.Println("Hello Andryan")
	} else if name == "Ariyan" {
		fmt.Println("Hello Ariyan")
	} else {
		fmt.Println("Hi, siapa kamu?")
	}

	// Short statement
	if length := len(name); length > 7 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}