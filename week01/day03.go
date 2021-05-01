package main

import (
	"fmt"
	"strings"
)

// 思路：记录最长字串，每一次判断后面的元素是否在前面字串中，如果在则将字串截断，否则加入字串。当某一个字串大于之前的字串时，发生字串替换，最后计算最长字串长度
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := ""
	str := ""
	for _, item := range s{
		loc := strings.Index(str, string(item))  // 查看字串中这个元素的位置
		if loc >= 0{ // 有这个元素
			str = str[loc+1:] + string(item)
		} else {
			str = str + string(item)
		}

		if len(str) > len(max) {
			max = str
		}
	}

	return len(max)
}

func main()  {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}