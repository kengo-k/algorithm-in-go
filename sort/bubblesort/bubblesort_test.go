package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
		{[]int{5, 1, 4, 2, 8}, []int{1, 2, 4, 5, 8}},
		{[]int{}, []int{}}, // 空のリスト
	}

	for _, tt := range tests {
		BubbleSort(tt.input)
		if !reflect.DeepEqual(tt.input, tt.output) {
			t.Errorf("bubbleSort(%v) = %v, want %v", tt.input, tt.input, tt.output)
		}
	}
}
