package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1767A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1767A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1767A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1767A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1767A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1767A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1767A
Date: 12/16/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1767/A
Problem: CF1767A
**/
func (in *CF1767A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		in.NextString()
		a, b := in.NextInt(), in.NextInt()
		c, d := in.NextInt(), in.NextInt()
		e, f := in.NextInt(), in.NextInt()

		if (a == c || a == e || c == e) && (b == d || b == f || d == f) {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

	}
}

func NewCF1767A(r *bufio.Reader) *CF1767A {
	return &CF1767A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1767A(bufio.NewReader(os.Stdin)).Run()
}
