package mergesort

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	for ; l < len(left); l++ {
		result = append(result, left[l])
	}

	for ; r < len(right); r++ {
		result = append(result, right[r])
	}

	return result
}
