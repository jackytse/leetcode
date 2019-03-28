package main

import "fmt"

func main() {
	var nums = []int{1, 3, 4, 2}
	var target = 6
	var result = twoSum2(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		v, ok := m[complement]
		if ok && i != v {
			return []int{i, v}
		}
		m[nums[i]] = i
	}
	return nil
}
