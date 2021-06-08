package main

import "fmt"

//0x1002a48c0
//0x1002a48c8
//0x1002a48d0
//0x1002a48e0
//0x1002a48e8
//0x1002a48f0

// 1.二维数组在地址空间上是连续的
// 2.int 在 64 位 arm 上占 8 个字节
var array = [][]int{
	{1, 2, 3},
	{3, 4, 5},
}

func main() {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Printf("%p\n", &array[i][j])
		}
	}
}
