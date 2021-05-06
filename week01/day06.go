package main

import "math"

// https://leetcode-cn.com/problems/reverse-integer/
func reverse(x int) int {
	var rev int

	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}

		tmp := x%10 + rev*10
		rev = tmp
		x = x / 10
	}

	return rev
}

func main() {
	// (3 + 0) + (2 + 30) + (1 + 320) = -321
	println(reverse(-123))
}
