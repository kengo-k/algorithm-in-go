package bubblesort

//  1. Start from the beginning of the list and compare two adjacent elements.
//  2. If the element on the left is greater than the one on the right, swap them.
//  3. Continue this process to the end of the list. By the end of a single pass,
//     the largest element is guaranteed to have moved to the last position.
//  4. Treat the last element of the list as sorted, and repeat the process for the list up to the second last element.
//  5. Continue this process until the entire list is sorted.
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
