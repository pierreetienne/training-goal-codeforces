package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1438A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1438A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1438A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1438A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1438A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1438A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1438A
Date: 17/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1438/A
Problem: CF1438A
**/
func (in *CF1438A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()

		var sb strings.Builder
		for i, j := 0, 3; i < n; i++ {
			sb.WriteString(fmt.Sprintf("%d ", j))
		}
		fmt.Println(sb.String())
	}
}

func NewCF1438A(r *bufio.Reader) *CF1438A {
	return &CF1438A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1438A(bufio.NewReader(os.Stdin)).Run()
}
