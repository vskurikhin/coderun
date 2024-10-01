package graph

type Color int

const (
	_ Color = iota
	White
	Gray
	Black
)

type Edge interface {
	From() int
	To() int
}

type Graph[V Vertex] interface {
	AddEdge(u, v V)
	All(func(V) bool, func(V) bool)
	GetColors() []Color
}

type Vertex interface {
	GetIndex() int
	GetLabel() string
	GetValue() interface{}
}
