package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF361A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF361A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF361A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF361A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF361A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF361A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF361A
Date: 5/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/361/A
Problem: CF361A
**/
func (in *CF361A) Run() {
	n, k := in.NextInt(), in.NextInt()
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		m[i][i] = k
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(m[i][j], " ")
		}
		fmt.Println()
	}
}

func NewCF361A(r *bufio.Reader) *CF361A {
	return &CF361A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF361A(bufio.NewReader(os.Stdin)).Run()
}
