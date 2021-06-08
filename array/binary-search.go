package main

import "fmt"

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		middle := lo + (hi-lo)>>2

		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			lo = middle + 1
		} else if nums[middle] > target {
			hi = middle - 1
		}
	}

	return -1
}

func main() {
	index := search([]int{-1, 0, 3, 5, 9, 12}, 2)

	fmt.Println(index)
}
