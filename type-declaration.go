package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool
	
	var ktpAriyan NoKTP = "1234567890"
	var marriedStatus Married = false
	fmt.Println(ktpAriyan, marriedStatus)
}