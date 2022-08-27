package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1315A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1315A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1315A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1315A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1315A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1315A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1315A
Date: 8/25/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1315/A
Problem: CF1315A
**/
func (in *CF1315A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, b, x, y := in.NextInt(), in.NextInt(), in.NextInt()+1, in.NextInt()+1

		max := Max((b-y)*a, Max((a-x)*b, Max(b*(x-1), a*(y-1))))
		fmt.Println(max)
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewCF1315A(r *bufio.Reader) *CF1315A {
	return &CF1315A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1315A(bufio.NewReader(os.Stdin)).Run()
}
