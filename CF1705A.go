package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1705A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1705A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1705A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1705A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1705A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1705A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1705A
Date: 20/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1705/A
Problem: CF1705A
**/
func (in *CF1705A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, x := in.NextInt(), in.NextInt()
		arr := make([]int, 2*n)
		for i := 0; i < len(arr); i++ {
			arr[i] = in.NextInt()
		}

		sort.Ints(arr)
		ans := true
		for i := 0; i < n; i++ {
			if arr[i+n]-arr[i] < x {
				ans = false
				break
			}
		}

		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1705A(r *bufio.Reader) *CF1705A {
	return &CF1705A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1705A(bufio.NewReader(os.Stdin)).Run()
}
