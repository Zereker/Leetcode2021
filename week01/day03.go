package main

/**
 * @Description:参考了题解，采用滑动窗口的思想解决
 * @link: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/hua-jie-suan-fa-3-wu-zhong-fu-zi-fu-de-zui-chang-z/
 */
func lengthOfLongestSubstring(s string) int {
	begin, result := 0, 0
	//temp用来存储字符和下标的关系
	tempMap := make(map[byte]int)
	sByte := []byte(s)
	for end := 0; end < len(sByte); end++ {
		//若map中已存在相同字符，则将begin更新到 最后这个字符出现的较大位置处
		if value, ok := tempMap[sByte[end]]; ok {
			begin = max(begin, value)
		}
		//结果等于尾-头+1的最大值
		result = max(result, end-begin+1)
		//更新字符的最后一个下标
		tempMap[sByte[end]] = end + 1
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
