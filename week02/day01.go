package main

/*
	https://leetcode-cn.com/problems/regular-expression-matching/
*/

// 递归回溯
func isMatch(s string, p string) bool {
	// 在 p 为空后， 只需要查看 s 是否为空即可
	if len(p) == 0 {
		return len(s) == 0
	}

	// 查看首元素是否一致
	firstMatch := len(s) != 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' { // 如果下一个字符是 '*'
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else { // 一般情况
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

// 动态规划
// 找到最优子结构
func isMatch2(s string, p string) bool {
	firstMatch := func(s, p string, i, j int) bool {
		return s[i] == p[j] || p[j] == '.'
	}

	// 要处理空字符串的情况，
	// 比如说 p 为空, s 为空时的情况；或者 p 为空, s 不为空时的情况
	m := len(s) + 1
	n := len(p) + 1

	// p 的前 [n-1] 个字符能否匹配 s 的前 [m-1] 个字符
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}

	dp[0][0] = true
	for j := 2; j <= len(p); j++ {
		dp[0][j] = p[j-1] == '*' && dp[0][j-2]
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if p[j] == '*' {
				dp[i+1][j+1] = dp[i+1][j-1] || firstMatch(s, p, i, j-1) && dp[i][j+1]
			} else {
				dp[i+1][j+1] = firstMatch(s, p, i, j) && dp[i][j]
			}
		}
	}

	return dp[len(s)][len(p)]
}

func main() {
	println(isMatch2("mississippi", "mis*is*p*."))
}
