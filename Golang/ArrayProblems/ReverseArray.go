package main

import "fmt"

func main() {
	fmt.Print(reverseArray([]int{1, 2, 3}))
}

func reverseArray(arr1 []int) []int {
	left := 0
	right := len(arr1) - 1
	for left <= right {
		arr1[left], arr1[right] = arr1[right], arr1[left]
		left++
		right--
	}
	return arr1
}
