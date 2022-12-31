package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1695A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1695A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1695A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1695A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1695A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1695A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1695A
Date: 9/16/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1695/A
Problem: CF1695A
**/
func (in *CF1695A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt(), in.NextInt()
		max := -1000000001
		x, y := -1, -1
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				val := in.NextInt()
				if val > max {
					max = val
					x = i
					y = j
				}
			}
		}

		fmt.Println(maxi(n-x, x+1) * maxi(m-y, y+1))
	}
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewCF1695A(r *bufio.Reader) *CF1695A {
	return &CF1695A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1695A(bufio.NewReader(os.Stdin)).Run()
}
