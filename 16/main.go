package main

import "fmt"

func main() {
	arr := []int{9, 5, 2, 7, 1, 8, 3, 6, 4}

	fmt.Println("Before sorting:", arr)

	quicksort(arr)

	fmt.Println("After sorting:", arr)
}

func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivotIndex := partition(arr)

	quicksort(arr[:pivotIndex])
	quicksort(arr[pivotIndex+1:])
}

func partition(arr []int) int {
	pivotIndex := len(arr) - 1
	pivot := arr[pivotIndex]

	i := 0
	for j := 0; j < pivotIndex; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[pivotIndex] = arr[pivotIndex], arr[i]

	return i
}
