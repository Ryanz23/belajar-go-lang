package main

import "fmt"

func main() {
	names := [...]string{"Ariyan", "Andryan", "Aryja", "Bekasi", "Jawa Barat", "Indonesia"}
	slice := names[2:5]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// names[3] = "Bandung"
	// fmt.Println(slice)

	// slice[0] = "Ganti Nama"
	// fmt.Println(names)

	// modifikasi slice
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	weekend := days[5:]
	weekend[0] = "Sabtu Ceria"
	weekend[1] = "Minggu Ceria"
	fmt.Println(days) // [senin selasa rabu kamis jumat Sabtu Ceria Minggu Ceria]

	weekend2 := append(weekend, "Libur Lagi")
	weekend2[0] = "Sudah Waktunya"
	fmt.Println(weekend2) // [Sudah Waktunya Minggu Ceria Libur Lagi]
	fmt.Println(days)     // [senin selasa rabu kamis jumat Sabtu Ceria Minggu Ceria]

	// membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ariyan"
	newSlice[1] = "Andryan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice
	copySlice :=(make([]string, len(newSlice), cap(newSlice)))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// perbedaan array dan slice
	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}