package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF34B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF34B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF34B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF34B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF34B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF34B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF34B
Date: 8/20/2022
User: wotan
URL: https://codeforces.com/problemset/problem/34/B
Problem: CF34B
**/
func (in *CF34B) Run() {
	n, m := in.NextInt(), in.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
	}
	sort.Ints(arr)
	ans := 0
	for i := 0; i < n; i++ {
		if m > 0 {
			if arr[i] < 0 {
				ans += -arr[i]
			}
			m--
		}
	}
	fmt.Println(ans)
}

func NewCF34B(r *bufio.Reader) *CF34B {
	return &CF34B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF34B(bufio.NewReader(os.Stdin)).Run()
}
