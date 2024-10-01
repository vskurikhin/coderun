package graph_bfs

import "coderun/v1/graph"

type GraphAdjacencyList struct {
	adjacency [][]int
	color     []graph.Color
	distance  []int
	previous  []int
}

func New(capacity int) GraphAdjacencyList {
	return GraphAdjacencyList{
		adjacency: make([][]int, capacity),
		color:     make([]graph.Color, capacity),
		distance:  make([]int, capacity),
		previous:  make([]int, capacity),
	}
}

func (g *GraphAdjacencyList) AddEdge(u, v int) {
	g.growAdjacency(u)
	g.growColor(u)
	g.growDistance(u)
	g.growPrevious(u)
	g.adjacency[u] = append(g.adjacency[u], v)
}

func (g *GraphAdjacencyList) BFS(s int) {
	planned := []int{s}
	g.color[s] = graph.Gray
	g.distance[s] = 0
	for len(planned) > 0 {
		u := planned[0]
		planned = planned[1:]
		for _, v := range g.adjacency[u] {
			if g.color[v] == 0 {
				g.distance[v] = g.distance[u] + 1
				g.previous[v] = u
				g.color[v] = graph.Gray
				planned = append(planned, v)
			}
		}
		g.color[u] = graph.Black
	}
}

func (g *GraphAdjacencyList) ShortestPath(v int) []int {
	var path []int
	currentVertex := v
	for currentVertex != 0 {
		path = append(path, currentVertex)
		currentVertex = g.previous[currentVertex]
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func (g *GraphAdjacencyList) growAdjacency(u int) {
	for i := len(g.adjacency); i < (u + 1); i++ {
		g.adjacency = append(g.adjacency, []int{})
	}
}

func (g *GraphAdjacencyList) growColor(u int) {
	for i := len(g.color); i < (u + 1); i++ {
		g.color = append(g.color, 0)
	}
}

func (g *GraphAdjacencyList) growPrevious(u int) {
	for i := len(g.previous); i < (u + 1); i++ {
		g.previous = append(g.previous, 0)
	}
}

func (g *GraphAdjacencyList) growDistance(u int) {
	for i := len(g.distance); i < (u + 1); i++ {
		g.distance = append(g.distance, 0)
	}
}
