package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF609A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF609A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF609A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF609A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF609A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF609A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF609A
Date: 25/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/609/A
Problem: CF609A
**/
func (in *CF609A) Run() {
	n := in.NextInt()
	m := in.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
	}
	sort.Ints(arr)
	ans := 1
	for i := n - 1; i >= 0; i-- {
		m -= arr[i]
		if m <= 0 {
			break
		}
		ans++
	}
	fmt.Println(ans)
}

func NewCF609A(r *bufio.Reader) *CF609A {
	return &CF609A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF609A(bufio.NewReader(os.Stdin)).Run()
}
