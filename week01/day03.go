package main

/**
 * @Description:双重for循环暴力破解
 */
func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	temp := []byte(s)
	result, begin := 0, 0
	for begin < len(temp) {
		tempMap := make(map[byte]bool)
		tempLength := 0
		for i := begin; i < len(temp); i++ {
			if _, ok := tempMap[temp[i]]; !ok {
				tempMap[temp[i]] = true
				tempLength++
			} else {
				break
			}
			result = max(tempLength, result)
		}
		begin++
	}

	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
