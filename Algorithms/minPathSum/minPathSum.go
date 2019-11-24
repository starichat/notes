// 矩阵的最小路径和
package main

import "fmt"

func main() {

	myArray := [][]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}
	fmt.Println(minPathSum(myArray))

}

func minPathSum(arr [][]int) []int {

	if arr == nil {
		return nil
	}
	var result []int
	var less int
	//minPathSum = arr[0][0]
	// 引入两个坐标，一个向右，一个向右，获取向右或者向下的数值
	// 取其中最小值，并记录执行的坐标，递归执行将该坐标即为起始点。
	i := 0
	j := 0
	for j < len(arr) && i < len(arr[0]) {
		flag := arr[i][j+1] < arr[i+1][j]
		fmt.Println(arr[i][j+1], arr[i+1][j])
		fmt.Println(flag)

		if flag {
			less = arr[i][j+1]
			j++
		} else {
			less = arr[i+1][j]
			i++
		}
		fmt.Println(less)
		fmt.Println("i,j", i, j)
		result = append(result, less)
	}

	return result
}
