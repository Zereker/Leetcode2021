package main

func searchInsert(nums []int, target int) int {
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

	if hi <= 0 {
		return 0
	}

	if lo > len(nums)-1 {
		return len(nums)
	}

	panic(nil)
}

func main() {
	//println(searchInsert([]int{1, 3, 5, 6}, 5))
	println(searchInsert([]int{1, 3, 5, 6}, 2))
	//println(searchInsert([]int{1, 3, 5, 6}, 7))
	//println(searchInsert([]int{1, 3, 5, 6}, 0))
}
