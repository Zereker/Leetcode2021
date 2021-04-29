package main

import "fmt"

/*
	https://leetcode-cn.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	if len(nums) <= 0 {
		return nil
	}

	numMap := make(map[int]int, len(nums)) // key: 元素， value: index

	for index, item:= range nums{
		if _, ok := numMap[target-item]; ok {
			return []int{index, numMap[target-item]}
		} else {
			numMap[item] = item
		}
	}
	return nil
}

func main() {
	res := twoSum([]int{2,7,11,15}, 9)
	fmt.Println(res)
}