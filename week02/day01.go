package main

/*
	https://leetcode-cn.com/problems/regular-expression-matching/
*/
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		}

		return false
	}

	var (
		sIndex  int
		anyByte byte
	)

	// 使用 p 做匹配会简单点
	for i := 0; i < len(p); i++ {
		if sIndex > len(s)-1 { // p 的数量都大于了 s 的数量，明显错误
			return false
		}

		pChar := p[i]
		sChar := s[sIndex]

		if pChar == '*' {
			if anyByte == '.' { // 第一个字符为空，p 怎么匹配
				return false
			}

		} else if pChar == '.' {
			if sChar < 'a' || sChar > 'z' {
				return false
			}
		} else {
			if pChar != sChar {
				return false
			}
		}

		sIndex++
	}

	return true
}

func main() {
	println(isMatch("mississippi", "mis*is*p*."))
}
