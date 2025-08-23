package main

import "fmt"

func main() {
	name := "Ariyan"

	switch name {
	case "Ariyan":
		fmt.Println("Hello Ariyan")
	case "Andryan":
		fmt.Println("Hello Andryan")
	case "Bekasi":
		fmt.Println("Welcome to Bekasi")
	default:
		fmt.Println("Kenalan dong")
	}

	// Short statement
	// switch length := len(name); length > 10 {
	// case true:
	// 	fmt.Println("Terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length < 3:
		fmt.Println("Nama terlalu pendek")
	default:
		fmt.Println("Nama sudah benar")
	}
}
