package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	var conver func(nums []int) map[int]int

	conver = func(nums []int) map[int]int {
		result := make(map[int]int, 0)

		for _, num := range nums {
			if _, ok := result[num]; ok {
				result[num] += 1
			} else {
				result[num] = 1
			}
		}

		return result
	}

	m1 := conver(nums1)
	m2 := conver(nums2)

	var counts []int
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok { // 说明没有出现交集
			continue
		}

		// 出现了交集
		// 以出现次数最少的进行遍历插入
		if v1 > v2 { //
			for i := 0; i < v2; i++ {
				counts = append(counts, k)
			}
		} else {
			for i := 0; i < v1; i++ {
				counts = append(counts, k)
			}
		}
	}

	return counts
}

func main() {
	ints := intersect([]int{1, 2, 2, 1}, []int{2, 2})
	ints = intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	fmt.Printf("%+v", ints)
}
