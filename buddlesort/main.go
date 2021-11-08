/**
冒泡排序
 */

package main

import "fmt"

func BubbleSort(list []int) {
	n := len(list)
	didSwap := false

	for i := n-1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}

		if !didSwap {
			return
		}
	}
}

func main()  {
	sortData := []int{10, 5, 2, 7, 3, 8}
	BubbleSort(sortData)
	fmt.Println(sortData)

	sortData1 := []int{1,2,3,4,5}
	BubbleSort(sortData1)
	fmt.Print(sortData1)
}