package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke: ", counter)
	}

	slice := []string{"Ariyan", "Andryan", "Aryja"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	names := []string{"Ariyan", "Andryan", "Aryja"}
	for i, name := range names {
		fmt.Println("Index", i, "=", name)
		// fmt.Println(name)
	}

	person := make(map[string]string)
	person["name"] = "Ariyan"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}