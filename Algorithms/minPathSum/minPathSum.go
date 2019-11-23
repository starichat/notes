// 矩阵的最小路径和
package main

func main() {

}

func minPathSum(arr [][]int) {
	minPathSum := 0
	if arr == nil {
		return nil
	}
	minPathSum = arr[0][0]
	// 引入两个坐标，一个向右，一个向右，获取向右或者向下的数值
	// 取其中最小值，并记录执行的坐标，递归执行将该坐标即为起始点。

}
