/*
 * This file was last modified at 2024.01.04 15:07 by Victor N. Skurikhin.
 * main.go
 * $Id$
 *
 * Рассмотрим три числа aa, bb и cc. Упорядочим их по возрастанию.
 * Какое число будет стоять между двумя другими?
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

func flush()                   { _ = out.Flush() }
func scan(a ...interface{})    { _, _ = fmt.Fscan(in, a...) }
func printLn(a ...interface{}) { _, _ = fmt.Fprintln(out, a...) }

func do() {
	a := make([]int, 3)
	for i := 0; i < 3; i++ {
		scan(&a[i])
	}
	slices.Sort(a)
	printLn(a[1])
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
