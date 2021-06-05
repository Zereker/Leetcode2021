package main

func longestPalindrome(s string) string {
	size := len(s)

	if size == 1 {
		return s
	}

	var result string // 最长回文子串

	dp := make([][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, size)
	}

	for i := size - 1; i >= 0; i-- {
		for j := i; j < size; j++ {
			if i == j {
				dp[i][j] = true
			} else if j == i+1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
		}
	}

	return result
}
