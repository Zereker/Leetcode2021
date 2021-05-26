package main

import "fmt"

func backtracking(n int, start int, k int, path *[]int, result *[][]int) {
	if len(*path) == k {
		tmp := make([]int, len(*path))
		copy(tmp, *path)

		*result = append(*result, tmp)
		return
	}

	for i := start; i <= n-(k-len(*path))+1; i++ {
		*path = append(*path, i)

		backtracking(n, i+1, k, path, result)

		*path = (*path)[:len(*path)-1]
	}
}

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)

	backtracking(n, 1, k, &path, &result)

	return result
}

func main() {
	result := combine(4, 2)
	fmt.Printf("%+v\n", result)
}
