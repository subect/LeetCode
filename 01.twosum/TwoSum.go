package main

import "fmt"

func TwoSum(slice []int, target int) []int {
	currMap := make(map[int]int)
	for index, value := range slice {
		another := target - value
		if _, ok := currMap[another]; ok {
			return []int{currMap[another], index}
		}
		currMap[value] = index
	}
	return nil
}

func main() {
	slice := []int{3, 2, 4}
	result := TwoSum(slice, 6)
	fmt.Println(result)
}
