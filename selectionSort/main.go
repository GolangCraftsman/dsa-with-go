package main

import "fmt"

func main() {
	arr := []int{5, 4, 1, 2, 3}
	// arr := []int{1, 2, 3, 4, 5}

	fmt.Println("Before sort: ", arr)
	selectionSort(arr)
	fmt.Println("After sort: ", arr)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
