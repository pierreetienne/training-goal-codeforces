package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1554A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1554A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1554A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1554A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1554A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1554A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1554A
Date: 8/11/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1554/A
Problem: CF1554A
**/
func (in *CF1554A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		arr := make([]int64, n)
		sol := int64(0)
		for i := 0; i < n; i++ {
    		arr[i] = in.NextInt64()
    		if i > 0 {
				sol = Max(sol, arr[i]*arr[i-1])
			}
		}
		fmt.Println(sol)
	}
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func NewCF1554A(r *bufio.Reader) *CF1554A {
	return &CF1554A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1554A(bufio.NewReader(os.Stdin)).Run()
}
