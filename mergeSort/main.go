package main

import "fmt"

func main() {
	arr := []int{5, 4, 1, 2, 3}
	// arr := []int{1, 2, 3, 4, 5}

	fmt.Println("Before sort: ", arr)
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println("After sort: ", arr)
}

func mergeSort(arr []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		mergeSort(arr, low, mid)
		mergeSort(arr, mid+1, high)
		merge(arr, mid, low, high)
	}
}

func merge(arr []int, mid, low, high int) {
	// arr = []int{1, 3, 5, 2, 4, 6, 8}
	sortArr := make([]int, len(arr))
	i := low
	j := mid + 1
	k := low
	for i <= mid && j <= high {
		if arr[i] < arr[j] {
			sortArr[k] = arr[i]
			i++
		} else {
			sortArr[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		sortArr[k] = arr[i]
		k++
		i++
	}

	for j <= high {
		sortArr[k] = arr[j]
		k++
		j++
	}

	for i := low; i <= high; i++ {
		arr[i] = sortArr[i]
	}

	fmt.Println("F=", sortArr)
	fmt.Println("A=", arr)
}
