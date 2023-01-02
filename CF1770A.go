package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1770A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1770A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1770A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1770A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1770A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1770A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1770A
Date: 12/30/2022
User: wotan
URL: https://codeforces.com/contests/1770/A
Problem: CF1770A
**/
func (in *CF1770A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt(), in.NextInt()
		a := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = in.NextInt()
		}
		b := make([]int, m)
		for i := 0; i < m; i++ {
			b[i] = in.NextInt()
		}

		for i := 0; i < m; i++ {
			sort.Ints(a)
			a[0] = b[i]
		}

		sum := int64(0)
		for i := 0; i < n; i++ {
			sum += int64(a[i])
		}

		fmt.Println(sum)

	}
}

func NewCF1770A(r *bufio.Reader) *CF1770A {
	return &CF1770A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1770A(bufio.NewReader(os.Stdin)).Run()
}
