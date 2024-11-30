package main

import "fmt"

// var arr [n]T
// arr :=[n]T{v1,v2,...,vn}

func main() {
	var arr [3]int
	fmt.Println(arr)

	var arr1 = [3]int{10, 20, 30}
	fmt.Println(arr1)

	var arr3 = [...]int{10, 20, 30}
	fmt.Println(arr3[0])
	fmt.Println(arr3[2])

	var seasons [4]string
	seasons[0] = "spring"
	seasons[1] = "summer"
	seasons[2] = "fall"
	seasons[3] = "winter"
	fmt.Println(seasons)

}
