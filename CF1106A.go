package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1106A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1106A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1106A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1106A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1106A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1106A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1106A
Date: 3/12/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1106/A
Problem: CF1106A
**/
func (in *CF1106A) Run() {
	n := in.NextInt()
	m := make([][]uint8, n)

	for i := 0; i < n; i++ {
		x := make([]uint8, n)
		str := in.NextString()
		for j := 0; j < n; j++ {
			x[j] = str[j]
		}
		m[i] = x
	}

	dx := []int{0, -1, -1, 1, 1}
	dy := []int{0, -1, 1, -1, 1}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m[i][j] == 'X' {
				e := 0
				for p := 0; p < len(dx); p++ {
					x := dx[p] + i
					y := dy[p] + j
					if x >= 0 && y >= 0 && x < n && y < n && m[x][y] == 'X' {
						e++
					}
				}
				if e == 5 {
					count++
				}

			}

		}
	}

	fmt.Println(count)

}

func NewCF1106A(r *bufio.Reader) *CF1106A {
	return &CF1106A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1106A(bufio.NewReader(os.Stdin)).Run()
}
