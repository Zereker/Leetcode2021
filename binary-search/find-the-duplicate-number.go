package main

func findDuplicate(nums []int) int {
	var (
		// 这里是数值而非索引
		lo = 1
		hi = len(nums) - 1
	)

	for lo <= hi {
		middle := lo + ((hi - lo) >> 1)
		counter := 0

		for i := 0; i < len(nums); i++ {
			if nums[i] <= middle {
				counter++
			}
		}

		if counter <= middle {
			lo = middle + 1
		} else {
			hi = middle - 1

		}
	}

	return lo

}
