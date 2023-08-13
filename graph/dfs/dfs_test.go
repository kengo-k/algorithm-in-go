package dfs

import (
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
	g := New()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	expected := []int{0, 1, 2, 3}
	if result := g.DFS(0); !reflect.DeepEqual(result, expected) {
		t.Errorf("DFS from vertex 0, got: %v, want: %v", result, expected)
	}
}

/*
//     0
//    / \
//   v   v
//   1-->2
//       |
//       v
//       3
*/
func TestPathExists(t *testing.T) {
	g := New()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)

	tests := []struct {
		src, dest int
		want      bool
	}{
		{0, 1, true},
		{0, 3, true},
		{3, 0, false}, // No backward path from 3 to 0 in this directed graph.
		{1, 3, true},
	}

	for _, tt := range tests {
		got := g.PathExists(tt.src, tt.dest)
		if got != tt.want {
			t.Errorf("PathExists(%d, %d) = %v; want %v", tt.src, tt.dest, got, tt.want)
		}
	}
}
