package graph_bfs

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

func flush()                   { _ = out.Flush() }
func scan(a ...interface{})    { _, _ = fmt.Fscan(in, a...) }
func printLn(a ...interface{}) { _, _ = fmt.Fprintln(out, a...) }

func do() {
	var n, m int
	scan(&n, &m)
	g := New(n)
	for i := 0; i < m; i++ {
		var u, v int
		scan(&u, &v)
		g.AddEdge(u, v)
	}
	g.BFS(1)
	printLn(g.ShortestPath(2))
	flush()
}

func wrap(i io.Reader, o io.Writer, do func()) {
	in = bufio.NewReader(i)
	out = bufio.NewWriter(o)
	do()
}

func TestFileDo(t *testing.T) {
	var tests = []struct {
		file  string
		name  string
		input string
		want  string
	}{
		{name: "Test case #01",
			file: "01.txt",
			want: "[1 4 2]\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.file != "" {
				buf, err := os.ReadFile(test.file)
				if err != nil {
					fmt.Print(err)
				}
				test.input = string(buf)
			}
			var b bytes.Buffer
			w := bufio.NewWriter(&b)
			wrap(strings.NewReader(test.input), w, do)
			got := b.String()
			assert.Equal(t, test.want, got)
		})
	}
}
