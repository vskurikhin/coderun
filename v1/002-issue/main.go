/*
 * This file was last modified at 2024.01.04 15:07 by Victor N. Skurikhin.
 * main.go
 * $Id$
 */

// В каждой клетке прямоугольной таблицы N×M записано некоторое число.
// Изначально игрок находится в левой верхней клетке.
// За один ход ему разрешается перемещаться в соседнюю клетку либо вправо,
// либо вниз (влево и вверх перемещаться запрещено).
// При проходе через клетку с игрока берут столько килограммов еды,
// какое число записано в этой клетке (еду берут также за первую и последнюю клетки его пути).
//
// Требуется найти минимальный вес еды в килограммах, отдав которую игрок может попасть в правый нижний угол.
//
// Формат ввода
// Вводятся два числа N и M — размеры таблицы (1 ≤ N ≤ 20, 1 ≤ M ≤ 20).
// Затем идет N строк по M чисел в каждой — размеры штрафов в килограммах
// за прохождение через соответствующие клетки (числа от 0 до 100).
//
// Формат вывода
// Выведите минимальный вес еды в килограммах, отдав которую можно попасть в правый нижний угол.

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

func flush()                   { _ = out.Flush() }
func scan(a ...interface{})    { _, _ = fmt.Fscan(in, a...) }
func printLn(a ...interface{}) { _, _ = fmt.Fprintln(out, a...) }

type Edge struct {
	from int
	to   int
	cost int
}

type Cell struct {
	number int
	value  int
}

func dijkstra(n int, edges []Edge) [][]int {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				adj[i][j] = 0
			} else {
				adj[i][j] = math.MaxInt32
			}
		}
	}
	for _, e := range edges {
		adj[e.from][e.to] = e.cost
	}
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		copy(dist[i], adj[i])
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	return dist
}

func do() {
	var n, m int
	scan(&n, &m)
	a := make([][]Cell, n)
	k := 0
	for i := 0; i < n; i++ {
		a[i] = make([]Cell, m)
		for j := 0; j < m; j++ {
			scan(&a[i][j].value)
			a[i][j].number = k
			k++
		}
	}
	e := make([]Edge, 0)
	for i, b := range a {
		for j, c := range b {
			if j < len(b)-1 {
				e = append(e, Edge{c.number, b[j+1].number, b[j+1].value})
			}
			if i < len(a)-1 {
				e = append(e, Edge{c.number, a[i+1][j].number, a[i+1][j].value})
			}
		}
	}
	b := a[len(a)-1]
	target := b[len(b)-1]
	result := dijkstra(len(e), e)
	if len(result) > 1 {
		printLn(a[0][0].value + result[0][target.number])
	} else {
		printLn(a[0][0].value)
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

/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
