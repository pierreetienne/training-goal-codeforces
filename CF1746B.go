package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1746B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1746B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1746B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1746B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1746B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1746B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1746B
Date: 10/19/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1746/B
Problem: CF1746B
**/
func (in *CF1746B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		ans := 0
		arr := make([]int, n)
		cpy := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			cpy[i] = arr[i]
		}

		sort.Ints(arr)
		for i := n - 1; i >= 0; i-- {
			if arr[i] == 1 && cpy[i] != 1 {
				ans++
			}
		}

		fmt.Println(ans)
	}
}

func NewCF1746B(r *bufio.Reader) *CF1746B {
	return &CF1746B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1746B(bufio.NewReader(os.Stdin)).Run()
}
