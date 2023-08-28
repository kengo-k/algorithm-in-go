package quicksort

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	fmt.Println("sort start: ", arr)
	left, right := 0, len(arr)-1

	pivotIndex := len(arr) / 2

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
	fmt.Printf("  swap pivot: arr=%v, pivot=%d\n", arr, arr[right])

	for i := range arr {
		fmt.Printf("  %d < %d => %v\n", arr[i], arr[right], arr[i] < arr[right])
		if arr[i] < arr[right] {
			fmt.Printf("  swap index: i=%d, left=%d\n", i, left)
			arr[i], arr[left] = arr[left], arr[i]
			left++
			fmt.Printf("  partition: arr=%v, left=%d\n", arr, left)
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr
}
