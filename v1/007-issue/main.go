/*
 * This file was last modified at 2024.01.04 15:07 by Victor N. Skurikhin.
 * main.go
 * $Id$
 */

// Дан неориентированный граф, возможно с петлями и кратными ребрами.
// Необходимо найти компоненту связности, содержащую вершину с номером 1.
//
// Формат ввода
// В первой строке записаны два целых числа N (1 ≤ N ≤ 10³) и M (0 ≤ M ≤ 5×10⁵) — количество вершин и ребер в графе.
// В последующих M строках перечислены ребра — пары чисел, определяющие номера вершин, которые соединяют ребра.
//
// Вершины нумеруются с единицы.
//
// Формат вывода
// В первой строке выведите число KK — количество вершин в компоненте связности.
//
// Во второй строке выведите KK целых чисел — вершины компоненты связности, перечисленные в порядке возрастания номеров.
// Примечание
//
// Петля в графе — это ребро, которое соединяет вершину с самой собой.
//
// Кратные рёбра в графе — это рёбра, которые соединяют одну и ту же пару вершин.
//
// Компонента связности в неориентированном графе — это подмножество вершин таких,
// что все вершины достижимы друг из друга.

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
	g.DFS(1)
	result := countingSort(g.components, n+1)
	printLn(len(g.components))
	for i, v := range result {
		s := " "
		if i == 0 {
			s = ""
		}
		print(s, v)
	}
	printLn()
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
	adjacency  []map[int]struct{}
	color      []Color
	components []int
}

func NewGraphInt(capacity int) *GraphIntAdjacencyList {
	components := make([]int, capacity)
	return &GraphIntAdjacencyList{
		adjacency:  make([]map[int]struct{}, capacity),
		color:      make([]Color, capacity),
		components: components[0:0],
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
	g.components = append(g.components, v)
}

func countingSort(array []int, k int) []int {
	countedValues := make([]int, k)
	for _, value := range array {
		countedValues[value]++
	}
	index := 0
	for value := 0; value < k; value++ {
		for amount := 0; amount < countedValues[value]; amount++ {
			array[index] = value
			index++
		}
	}
	return array
}

/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
