package main

// https://leetcode-cn.com/problems/roman-to-integer/
func romanToInt(s string) int {
	var (
		romans = map[rune]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		}

		romansIndex = map[rune]int{
			'I': 1,
			'V': 2,
			'X': 3,
			'L': 4,
			'C': 5,
			'D': 6,
			'M': 7,
		}
	)

	var (
		result int
		prev   rune
	)
	for _, char := range s {
		roman := romans[char]
		romanIndex := romansIndex[char]


	}
}
