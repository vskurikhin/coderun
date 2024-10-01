package graph_dfs

import "coderun/v1/graph"

type GraphAdjacencyList[V graph.Vertex] struct {
	adjacency [][]V
	color     []graph.Color
	vertexes  []V
}

type GraphIntAdjacencyList struct {
	adjacency [][]int
	color     []graph.Color
}

type Vertex struct {
	Index int
	Label string
	Value interface{}
}

var _ graph.Vertex = (*Vertex)(nil)
var _ graph.Graph[graph.Vertex] = (*GraphAdjacencyList[graph.Vertex])(nil)

func New(capacity int) *GraphAdjacencyList[Vertex] {
	return &GraphAdjacencyList[Vertex]{
		adjacency: make([][]Vertex, capacity),
		color:     make([]graph.Color, capacity),
		vertexes:  make([]Vertex, capacity),
	}
}

func NewGraphInt(capacity int) *GraphIntAdjacencyList {
	return &GraphIntAdjacencyList{
		adjacency: make([][]int, capacity),
		color:     make([]graph.Color, capacity),
	}
}

func (g *GraphAdjacencyList[V]) AddEdge(u, v V) {
	g.growAdjacency(u.GetIndex())
	g.growColor(u.GetIndex())
	g.growVertexes(u)
	g.growVertexes(v)
	g.adjacency[u.GetIndex()] = append(g.adjacency[u.GetIndex()], v)
	g.vertexes[u.GetIndex()] = u
	g.vertexes[v.GetIndex()] = v
}

func (g *GraphAdjacencyList[V]) All(in func(V) bool, out func(V) bool) {
	for i := 0; i < len(g.color); i++ {
		if g.color[i] == graph.White {
			g.DFS(g.vertexes[i], in, out)
		}
	}
}

func (g *GraphAdjacencyList[V]) GetColors() []graph.Color {
	result := make([]graph.Color, len(g.color))
	copy(result, g.color)
	return result
}

func (g *GraphAdjacencyList[V]) DFS(v V, in func(V) bool, out func(V) bool) {
	g.color[v.GetIndex()] = graph.Gray
	in(v)
	for _, w := range g.adjacency[v.GetIndex()] {
		if g.color[w.GetIndex()] == graph.White {
			g.DFS(w, in, out)
		}
	}
	g.color[v.GetIndex()] = graph.Black
	out(v)
}

func (g *GraphAdjacencyList[V]) growAdjacency(u int) {
	for i := len(g.adjacency); i < (u + 1); i++ {
		g.adjacency = append(g.adjacency, []V{})
	}
}

func (g *GraphAdjacencyList[V]) growColor(u int) {
	for i := len(g.color); i < (u + 1); i++ {
		g.color = append(g.color, graph.White)
	}
}

func (g *GraphAdjacencyList[V]) growVertexes(u V) {
	for i := len(g.vertexes); i < (u.GetIndex() + 1); i++ {
		g.vertexes = append(g.vertexes, u)
	}
}

func (g *GraphIntAdjacencyList) AddEdge(u, v int) {
	g.growAdjacency(u)
	g.growColor(u)
	g.adjacency[u] = append(g.adjacency[u], v)
}

func (g *GraphIntAdjacencyList) All() {
	for i := 0; i < len(g.color); i++ {
		if g.color[i] == graph.White {
			g.DFS(i)
		}
	}
}

func (g *GraphIntAdjacencyList) DFS(v int) {
	g.color[v] = graph.Gray
	for _, w := range g.adjacency[v] {
		if g.color[w] == graph.White {
			g.DFS(w)
		}
	}
	g.color[v] = graph.Black
}

func (g *GraphIntAdjacencyList) growAdjacency(u int) {
	for i := len(g.adjacency); i < (u + 1); i++ {
		g.adjacency = append(g.adjacency, []int{})
	}
}

func (g *GraphIntAdjacencyList) growColor(u int) {
	for i := len(g.color); i < (u + 1); i++ {
		g.color = append(g.color, graph.White)
	}
}

func (v Vertex) GetIndex() int {
	return v.Index
}

func (v Vertex) GetLabel() string {
	return v.Label
}

func (v Vertex) GetValue() interface{} {
	return v.Value
}
