/*
 * This file was last modified at 2024.01.04 15:07 by Victor N. Skurikhin.
 * main.go
 * $Id$
 */

// Дан неориентированный невзвешенный граф, состоящий из NN вершин и MM ребер.
// Необходимо посчитать количество его компонент связности и вывести их.
//
// Напомним:
// Компонента связности в неориентированном графе - это подмножество вершин,
// таких что все вершины достижимы друг из друга.
//
// Формат ввода
// Во входном файле записано два числа N и M (0 < N ≤ 100_000, 0 ≤ M ≤ 100_000).
// В следующих M строках записаны по два числа i и j (1 ≤ i, j ≤ N), которые означают,
// что вершины i и j соединены ребром.
//
// Формат вывода
// В первой строчке выходного файла выведите количество компонент связности.
// Далее выведите сами компоненты связности в следующем формате:
// в первой строке количество вершин в компоненте, во второй - сами вершины в произвольном порядке.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

func flush()                   { _ = out.Flush() }
func scan(a ...interface{})    { _, _ = fmt.Fscan(in, a...) }
func print(a ...interface{})   { _, _ = fmt.Fprint(out, a...) }
func printLn(a ...interface{}) { _, _ = fmt.Fprintln(out, a...) }

func do() {
	var n, m int
	scan(&n, &m)
	g := NewGraphInt(n + 1)
	for j := 0; j < m; j++ {
		var v, u int
		scan(&v, &u)
		g.AddEdge(v, u)
		g.AddEdge(u, v)
	}
	g.All()
	printLn(g.currentComponents)
	for _, component := range g.components {
		if len(component) > 0 {
			printLn(len(component))
			for i, v := range component {
				s := " "
				if i == 0 {
					s = ""
				}
				print(s, v)
			}
			printLn()
		}
	}
	flush()
}

func wrap(i io.Reader, o io.Writer, do func()) {
	in = bufio.NewReader(i)
	out = bufio.NewWriter(o)
	do()
}

func main() {
	wrap(os.Stdin, os.Stdout, do)
}

type Color int

const (
	White Color = iota
	Gray
	Black
)

type GraphIntAdjacencyList struct {
	adjacency         []map[int]struct{}
	color             []Color
	components        [][]int
	currentComponents int
}

func NewGraphInt(capacity int) *GraphIntAdjacencyList {
	return &GraphIntAdjacencyList{
		adjacency:  make([]map[int]struct{}, capacity),
		color:      make([]Color, capacity),
		components: make([][]int, capacity),
	}
}

func (g *GraphIntAdjacencyList) AddEdge(u, v int) {
	if g.adjacency[u] == nil {
		g.adjacency[u] = make(map[int]struct{})
	}
	g.adjacency[u][v] = struct{}{}
}

func (g *GraphIntAdjacencyList) DFS(v int) {
	g.color[v] = Gray
	for w, _ := range g.adjacency[v] {
		if g.color[w] == White {
			g.DFS(w)
		}
	}
	g.color[v] = Black
	g.components[g.currentComponents] = append(g.components[g.currentComponents], v)
}

func (g *GraphIntAdjacencyList) All() {
	for i := 1; i < len(g.color); i++ {
		if g.color[i] == White {
			g.DFS(i)
			g.currentComponents++
		}
	}
}

/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
