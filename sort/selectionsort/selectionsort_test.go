package selectionsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{3, 1, 8, 6, 2, 1, 10}
	expected := []int{1, 1, 2, 3, 6, 8, 10}
	SelectionSort(arr)
	assert.Equal(t, arr, expected)
}
