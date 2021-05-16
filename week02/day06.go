package main

/*
	https://leetcode-cn.com/problems/implement-strstr/
*/

func next(s string) []int {
	m := len(s)

	next := make([]int, m) // 前缀表
	for i, j := 1, 0; i < m; j++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}

		if s[i] == s[j] {
			j++
		}

		next[i] = j
	}

	return next
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	n := len(haystack)
	next := next(needle)

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

		if haystack[i] == needle[j] {
			j++
		}

		if j == len(needle) {
			return i - n + 1
		}
	}

	return -1
}
