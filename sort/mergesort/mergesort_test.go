package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	arr := []int{10, 6, 8, 5, 7, 3, 0, 2, 9, 1, 4}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	actual := MergeSort(arr)
	assert.Equal(t, actual, expected)
}
