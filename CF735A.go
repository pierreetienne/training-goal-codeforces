package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF735A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF735A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF735A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF735A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF735A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF735A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF735A
Date: 12/8/2022
User: wotan
URL: https://codeforces.com/problemset/problem/735/A
Problem: CF735A
**/
func (in *CF735A) Run() {
	n, k := in.NextInt(), in.NextInt()
	a := in.NextString()
	str := make([]uint8, n)
	g := 0
	t := 0

	for i := 0; i < n; i++ {
		str[i] = a[i]
		if a[i] == 'G' {
			g = i
		}
		if a[i] == 'T' {
			t = i
		}
	}

	if g > t {
		str[t], str[g] = str[g], str[t]
		t, g = g, t
	}

	ws := false
	for i := g; i <= t; i += k {
		if i == t {
			ws = true
		}

		if str[i] == '#' {
			break
		}

	}

	if ws {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func NewCF735A(r *bufio.Reader) *CF735A {
	return &CF735A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF735A(bufio.NewReader(os.Stdin)).Run()
}
