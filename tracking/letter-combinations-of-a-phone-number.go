package main

import "fmt"

func letterCombinations(digits string) []string {
	result := make([]string, 0)

	dict := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	var path string
	var backtracking func(total, curIndex int)

	backtracking = func(total, curIndex int) {
		if curIndex >= total {
			result = append(result, path)
			return
		}

		char := digits[curIndex]
		values := dict[char]

		for i := curIndex; i < len(values); i++ {
			path += string(b)

			backtracking(total, curIndex)

			path = path[:len(path)-1]
		}
	}

	backtracking(len(digits), 0)

	return result
}

func main() {
	result := letterCombinations("23")
	fmt.Printf("%+v", result)
}
