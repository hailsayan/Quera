package main

import "fmt"

func main() {
	var target int = 5
	sum := 0
	for i := 1; i <= target; i++ {
		sum += i
	}
	fmt.Println(sum)
}
