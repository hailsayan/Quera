package main

import (
	"fmt"
	"strconv"
)

func ConvertToDigitalFormat(hour, minute, second int) string {
	hh := strconv.Itoa(hour)
	mm := strconv.Itoa(minute)
	ss := strconv.Itoa(second)

	if hour < 10 {
		hh = "0" + hh
	}
	if minute < 10 {
		mm = "0" + mm
	}
	if second < 10 {
		ss = "0" + ss
	}

	return hh + ":" + mm + ":" + ss
}

func ExtractTimeUnits(seconds int) (int, int, int) {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	remainingSeconds := seconds % 60
	return hours, minutes, remainingSeconds
}

func main() {
	fmt.Println(ConvertToDigitalFormat(2, 23, 4))
	fmt.Println(ConvertToDigitalFormat(18, 3, 0))

	h, m, s := ExtractTimeUnits(5500)
	fmt.Printf("%d hours, %d minutes, %d seconds\n", h, m, s)
}
