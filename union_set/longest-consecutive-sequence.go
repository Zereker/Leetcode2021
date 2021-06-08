package main

import "fmt"

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})

	for _, v := range nums {
		set[v] = struct{}{}
	}

	var result int
	for k := range set {
		if _, ok := set[k-1]; ok {
			continue
		} else {
			var length int // len 记录以 num 为左边界的连续序列的长度

			for {
				if _, ok := set[k]; ok {
					length++
				} else {
					break
				}

				k = k + 1
			}

			if length > result {
				result = length
			}
		}
	}

	return result
}

func main() {
	consecutive := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	fmt.Println(consecutive)
}
