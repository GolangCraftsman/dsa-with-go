package main

import "fmt"

func main() {
	arr := []int{12, 13, 16, 13, 19}
	// arr := []int{1, 2, 3, 4, 5}

	fmt.Println("Before sort: ", arr)
	bubbleSort(arr)
	fmt.Println("After sort: ", arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("Pass: %d\n", i)
		isSorted := true
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				isSorted = false
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if isSorted {
			return
		}
	}
}
