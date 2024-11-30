package main

import "fmt"

func main() {
	arr := [...]int{10, 20, 30, 40, 50}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("element %d at index %d\n", arr[i], i)
	}
}
