package main

func search(nums []int, target int) int {
	var (
		lo = 0
		hi = len(nums) - 1
	)

	for lo <= hi {
		middle := lo + (hi-lo)/2
		value := nums[middle]

		if value == target {
			return middle
		} else if value < target {
			lo = middle + 1
		} else {
			hi = middle - 1
		}
	}

	// 未找到
	return -1
}

// 查找第一个或者最后一个值等于给定值的元素
func searchFirst(nums []int, target int) int {
	var (
		lo = 0
		hi = len(nums) - 1
	)

	for lo <= hi {
		middle := lo + (hi-lo)/2
		value := nums[middle]

		if value < target {
			lo = middle + 1
		} else if value > target {
			hi = middle - 1
		} else {
			if middle == 0 || nums[middle-1] != value {
				return middle
			} else {
				hi = middle - 1
			}
		}
	}

	// 未找到
	return -1
}

// 查找第一个大于等于给定值的元素
func searchMore(nums []int, target int) int {
	var (
		lo = 0
		hi = len(nums) - 1
	)

	for lo <= hi {
		middle := lo + (hi-lo)>>1
		value := nums[middle]

		if value >= target {
			if middle == 0 || nums[middle-1] < value {
				return middle
			} else {
				hi = middle - 1
			}
		} else if value < target {
			lo = middle + 1
		}
	}

	return -1
}

func main() {
	println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
