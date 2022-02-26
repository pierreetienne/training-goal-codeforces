package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF950A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF950A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF950A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF950A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF950A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF950A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF950A
Date: 24/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/950/A
Problem: CF950A
**/
func (in *CF950A) Run() {
	l, r, a := in.NextInt(), in.NextInt(), in.NextInt()

	mx := max(l, r)
	min := min(l, r)

	diff := mx - min

	if diff > 0 {
		if min+a <= mx {
			min += a
			a = 0
		} else {
			min = mx
			a -= diff
		}
	}
	if a == 0 {
		fmt.Println(min + min)
	} else {
		if a % 2 != 0 {
			a--
		}

		fmt.Println(min + mx + a)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NewCF950A(r *bufio.Reader) *CF950A {
	return &CF950A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF950A(bufio.NewReader(os.Stdin)).Run()
}
