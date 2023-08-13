package dfs

type Graph struct {
	AdjList map[int][]int
}

func New() *Graph {
	return &Graph{
		AdjList: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(v, w int) {
	g.AdjList[v] = append(g.AdjList[v], w)
}

func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	g.dfsUtil(start, visited, &result)
	return result
}

func (g *Graph) dfsUtil(v int, visited map[int]bool, result *[]int) {
	visited[v] = true
	*result = append(*result, v)

	for _, i := range g.AdjList[v] {
		if !visited[i] {
			g.dfsUtil(i, visited, result)
		}
	}
}
