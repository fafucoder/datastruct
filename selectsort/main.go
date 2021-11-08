/**
选择排序
 */

package main

import "fmt"

func SelectSort(list []int) {
	n := len(list)

	// 控制循环
	for i := 0; i < n; i++ {
		minIndex, minData := i, list[i]

		// 获取最小的数
		for j := i+1; j < n; j++ {
			if list[j] < minData {
				minIndex, minData = j, list[j]
			}
 		}

 		if minIndex != i {
 			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

func main()  {
	sortData := []int{10, 5, 2, 7, 3, 8}
	SelectSort(sortData)
	fmt.Println(sortData)

	sortData1 := []int{1,2,3,4,5}
	SelectSort(sortData1)
	fmt.Print(sortData1)
}