package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1680A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1680A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1680A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1680A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1680A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1680A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1680A
Date: 20/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1680/A
Problem: CF1680A
**/
func (in *CF1680A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		l1, r1, l2, r2 := in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()

		if r1 < l2  || r2 < l1{
			fmt.Println(l1 + l2)

		} else {
			fmt.Println(int(math.Max(float64(l1), float64(l2))))
		}

	}
}

func NewCF1680A(r *bufio.Reader) *CF1680A {
	return &CF1680A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1680A(bufio.NewReader(os.Stdin)).Run()
}
