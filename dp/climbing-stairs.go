package main

import "fmt"

func climbStairs(n int) int {
	// 1. 确定dp数组以及下标的含义
	// 2. 确定递推公式
	// 3. dp数组如何初始化
	// 4. 确定遍历顺序
	// 5. 举例推导dp数组
	dp := make([]int, 3)
	dp[1] = 1
	dp[2] = 2

	if n <= 2 {
		return dp[n]
	}

	for i := 3; i <= n; i++ {
		sum := dp[1] + dp[2]

		dp[1] = dp[2]
		dp[2] = sum
	}

	return dp[2]
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}
