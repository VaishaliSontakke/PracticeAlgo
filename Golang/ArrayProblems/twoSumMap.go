package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)
	for idx, val := range nums {
		diff := target - val
		if index1, ok := hmap[diff]; ok {
			return []int{index1, idx}
		}
		hmap[val] = idx
	}
	return nil

}
