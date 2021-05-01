package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	max := 0
	str := ""
	for _, item := range s{
		loc := strings.Index(str, string(item))  // 查看字串中这个元素的位置
		if loc >= 0{ // 有这个元素
			str = str[loc+1:] + string(item)
		} else {
			str = str + string(item)
		}

		if len(str) > max {
			max = len(str)
		}
	}

	return max
}

func main()  {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}