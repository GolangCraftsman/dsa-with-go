package main

import "fmt"

func main() {
	arr := []int{5, 4, 1, 2, 3}
	// arr := []int{1, 2, 3, 4, 5}

	fmt.Println("Before sort: ", arr)
	insertionSort(arr)
	fmt.Println("After sort: ", arr)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		key := arr[i]
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
