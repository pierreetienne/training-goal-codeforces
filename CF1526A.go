package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1526A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1526A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1526A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1526A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1526A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1526A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1526A
Date: 17/11/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1526/A
Problem: CF1526A
**/
func (in *CF1526A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt() * 2
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		sort.Ints(arr)
		var sol strings.Builder
		for p, i, j := 0, 0, n-1; p < n; p++ {
			if p%2 != 0 {
				sol.WriteString(fmt.Sprintf("%d ", arr[i]))
				i++
			} else {
				sol.WriteString(fmt.Sprintf("%d ", arr[j]))
				j--
			}
		}
		fmt.Println(sol.String())
	}
}

func NewCF1526A(r *bufio.Reader) *CF1526A {
	return &CF1526A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1526A(bufio.NewReader(os.Stdin)).Run()
}
