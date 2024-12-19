package main

import (
	"fmt"
	"math"
)

// FilterFunc تایپ تابعی که یک عدد صحیح را به عنوان ورودی می‌گیرد و یک بولین برمی‌گرداند.
type FilterFunc func(int) bool

// MapperFunc تایپ تابعی که یک عدد صحیح را به عنوان ورودی می‌گیرد و یک عدد صحیح برمی‌گرداند.
type MapperFunc func(int) int

// IsSquare بررسی می‌کند که آیا عدد ورودی مربع کامل است یا خیر.
func IsSquare(x int) bool {
	if x < 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(x)))
	return sqrt*sqrt == x
}

// IsPalindrome بررسی می‌کند که آیا عدد ورودی متقارن است یا خیر.
func IsPalindrome(x int) bool {
	str := fmt.Sprintf("%d", x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

// Abs قدر مطلق یک عدد صحیح را محاسبه می‌کند.
func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// Cube مکعب یک عدد صحیح را محاسبه می‌کند.
func Cube(num int) int {
	return num * num * num
}

// Filter عناصر ورودی را بر اساس شرط داده‌شده فیلتر می‌کند.
func Filter(input []int, f FilterFunc) []int {
	var result []int
	for _, v := range input {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map تابع داده‌شده را روی تک تک عناصر ورودی اعمال می‌کند و نتیجه را جایگزین می‌کند.
func Map(input []int, m MapperFunc) []int {
	result := make([]int, len(input))
	for i, v := range input {
		result[i] = m(v)
	}
	return result
}

func main() {
	// مثال برای Filter با تابع IsSquare
	fmt.Println(Filter([]int{1, 4, 10}, IsSquare)) // خروجی: [1 4]

	// مثال برای Filter با تابع IsPalindrome
	fmt.Println(Filter([]int{121, -121, 10}, IsPalindrome)) // خروجی: [121]

	// مثال برای Map با تابع Abs
	fmt.Println(Map([]int{1, -4, -10}, Abs)) // خروجی: [1 4 10]

	// مثال برای Map با تابع Cube
	fmt.Println(Map([]int{1, 2, 3}, Cube)) // خروجی: [1 8 27]
}
