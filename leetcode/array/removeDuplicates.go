package main

import (
	"fmt"
)

func main() {
	var nums []int = []int{1, 1, 2, 3, 3, 4}
	fmt.Println(nums)
	var newLen int = removeDuplicates(nums)
	fmt.Println(nums)
	fmt.Println(newLen)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	var index int = 1
	var last int = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != last {
			nums[index] = nums[i]
			index++
		}
		last = nums[i]
	}
	return index
}

func removeDuplicatesForNoSortedArray(nums []int) int {
	var index int = 0
	for i := 0; i < len(nums); i++ {
		var find bool = false
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				find = true
				break
			}
		}
		if !find {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func removeDuplicatesWithHash(nums []int) int {
	var m map[int]int = make(map[int]int)
	for _, v := range nums {
		m[v] = 1
	}

	var i int = 0
	for k, _ := range m {
		nums[i] = k
		i++
	}
	return i
}
