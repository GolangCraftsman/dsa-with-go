package main

import "fmt"

func main() {
	// arr := []int{9, 4, 4, 8, 7, 5, 6}
	// arr := []int{1, 2, 3, 4, 5}
	arr := []int{5, 4, 3, 2, 1}

	fmt.Println("Before sort: ", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("After sort: ", arr)
}

func quickSort(arr []int, low, high int) {
	var idx int

	if low < high {
		idx = partition(arr, low, high)
		quickSort(arr, low, idx-1)  //sort left array
		quickSort(arr, idx+1, high) //sort right array
	}
}

func partition(arr []int, low, high int) int {

	pivot := arr[low]
	i := low + 1
	j := high

	for {
		for arr[i] <= pivot && i < high {
			i++
		}
		for j > low && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			continue
		}
		break
	}
	arr[low], arr[j] = arr[j], arr[low]
	
	return j
}
