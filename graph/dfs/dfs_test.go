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
