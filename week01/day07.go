package main

import (
	"math"
	"strings"
)

// https://leetcode-cn.com/problems/string-to-integer-atoi/
func myAtoi(s string) int {
	var (
		sign   = 1
		result int
	)

	str := strings.TrimSpace(s)

	for index, b := range str {
		if b >= '0' && b <= '9' {
			result = result*10 + int(b-'0')
		} else if index == 0 && b == '+' {
			sign = 1
		} else if index == 0 && b == '-' {
			sign = -1
		} else {
			break
		}

		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
	}

	return sign * result
}
