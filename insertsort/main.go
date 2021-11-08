/**
 插入排序
 */
package main

import "fmt"

func InsertSort(list []int) {
	n := len(list)
	for i := 1; i <= n-1; i++ {
		j, deal:= i - 1, list[i]

		// 说明需要插入排序
		if deal < list[j] {
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j]
			}

			list[j+1] = deal
		}
	}
}

func main() {
	sortData := []int{10, 5, 2, 7, 3, 8}
	InsertSort(sortData)
	fmt.Println(sortData)

	sortData1 := []int{1,2,3,4,5}
	InsertSort(sortData1)
	fmt.Print(sortData1)
}
