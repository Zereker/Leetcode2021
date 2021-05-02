package main

/*
	https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

	思路：滑动窗口问题， 使用快，慢指针:
		当查询到重复元素时，慢指针 = 上次出现元素的位置+1 （注意：慢指针不能比当前慢指针还慢）
		否则：
			更新元素位置
			比较最长 子串的大小

*/

func lengthOfLongestSubstring(s string) int {
	var (
		slow, fast int // 快慢指针
		result     int // 最长子串的长度

		bytes = []byte(s)
		dict  [256]int
	)

	if len(s) == 0 {
		return 0
	}

	for i := range dict {
		dict[i] = -1
	}

	for fast = 0; fast < len(bytes); fast++ {
		curr := bytes[fast]

		if dict[curr] > -1 { // 出现重复元素
			slow = max(slow, dict[curr]+1) // 慢指针 = 上次出现的位置+1
		}

		dict[curr] = fast
		result = max(result, fast-slow+1) // 比较是不是最长子串
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
