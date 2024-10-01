/*
 * This file was last modified at 2024-03-03 10:26 by Victor N. Skurikhin.
 * main_test.go
 * $Id$
 */

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestFileDo(t *testing.T) {
	var tests = []struct {
		file  string
		name  string
		input string
		want  string
	}{
		{name: "Test case #01",
			file: "01.txt",
			want: "4\n1 2 3 4\n",
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
