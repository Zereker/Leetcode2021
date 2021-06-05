package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)

	var path []int

	// 回溯
	var backtracking func(target, sum int, index, k int)
	backtracking = func(target, sum int, index, k int) {
		if sum > target { // 剪枝操作
			return
		}

		if len(path) == k {
			if target == sum {
				result = append(result, append([]int{}, path...))
			}

			return
		}

		for i := index; i <= 9-(k-len(path))+1; i++ {
			sum += i
			path = append(path, i)

			backtracking(target, sum, i+1, k)

			sum -= i
			path = path[:len(path)-1]
		}
	}

	backtracking(n, 0, 1, k)
	return result
}

func main() {
	result := combinationSum3(2, 4)
	fmt.Printf("%+v\n", result)
}
