// 快排

package main

import "fmt"

func quickSort(list []int, begin, end int) {
	if begin < end {
		location := getCaption(list, begin, end)

		quickSort(list, 0, location -1)
		quickSort(list, location + 1, end)
	}
}

func getCaption(list []int, begin, end int) int {
	i, j := begin +1, end

	for i < j {
		if list[i] > list[begin] {
			list[i], list[j] = list[j], list[i]
			j--
		} else {
			i++
		}
	}

	if list[i] >= list[begin] {
		i--
	}

	list[begin], list[i] = list[i], list[begin]
	return i
}

func main() {
	sortData := []int{10, 5, 2, 7, 3, 8}
	quickSort(sortData, 0, len(sortData) - 1)
	fmt.Println(sortData)

	sortData1 := []int{1,2,3,4,5}
	quickSort(sortData1, 0, len(sortData1) - 1)
	fmt.Print(sortData1)
}
