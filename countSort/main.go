package main

import "fmt"

func main() {
	arr := []int{5, 4, 2, 1, 1}
	fmt.Println("Before sort: ", arr)
	countSort(arr)
	fmt.Println("After sort: ", arr)

}

func countSort(arr []int) {
	countArr := make([]int, getMax(arr)+1)

	for _, val := range arr {
		countArr[val]++
	}
	i, j := 0, 0

	for i < len(countArr) {
		if countArr[i] != 0 {
			arr[j] = i
			countArr[i]--
			j++
		} else {
			i++
		}
	}
}

func getMax(arr []int) (m int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] > m {
			m = arr[i]
		}
	}
	return
}
