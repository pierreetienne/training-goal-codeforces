package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1099A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1099A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1099A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1099A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1099A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1099A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1099A
Date: 12/5/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1099/A
Problem: CF1099A
**/
func (in *CF1099A) Run() {
	w, h := in.NextInt(), in.NextInt()
	u1, h1 := in.NextInt(), in.NextInt()
	u2, h2 := in.NextInt(), in.NextInt()

	cw := w
	ch := h
	for ch > 0 {
		cw += ch
		if h1 == ch {
			cw -= u1
		}
		if h2 == ch {
			cw -= u2
		}
		if cw < 0 {
			cw = 0
		}
		ch--
	}
	fmt.Println(cw)

}

func NewCF1099A(r *bufio.Reader) *CF1099A {
	return &CF1099A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1099A(bufio.NewReader(os.Stdin)).Run()
}
