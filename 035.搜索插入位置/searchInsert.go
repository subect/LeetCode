package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for index, values := range nums {
		if values >= target {
			return index
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 3, 5, 6}
	result := searchInsert(nums, 5)
	fmt.Println(result)
}
