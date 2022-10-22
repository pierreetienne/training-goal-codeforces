package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1749A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1749A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1749A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1749A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1749A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1749A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1749A
Date: 10/20/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1749/A
Problem: CF1749A
**/
func (in *CF1749A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt(), in.NextInt()
		cols := make([]bool, n)
		rows := make([]bool, n)

		for i := 0; i < m; i++ {
			x, y := in.NextInt(), in.NextInt()
			cols[x-1] = true
			rows[y-1] = true
		}

		c, r := false, false
		for i := 0; i < n; i++ {
			if !cols[i] {
				c = true
			}
			if !rows[i] {
				r = true
			}

		}
		if c && r {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func NewCF1749A(r *bufio.Reader) *CF1749A {
	return &CF1749A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1749A(bufio.NewReader(os.Stdin)).Run()
}
