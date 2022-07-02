package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1631A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1631A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1631A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1631A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1631A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1631A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1631A
Date: 4/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1631/A
Problem: CF1631A
**/
func (in *CF1631A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		a := make([]int, n)
		b := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = in.NextInt()
		}

		for i := 0; i < n; i++ {
			b[i] = in.NextInt()
		}

		sort.Ints(a)
		sort.Ints(b)

		for i, j := n-1, 0; i >= 0 && j < n; {
			if a[i] > b[j] {
				aux := a[i]
				a[i] = b[j]
				b[j] = aux
			}
			i--
			j++
		}

		sort.Ints(a)
		sort.Ints(b)

		fmt.Println(a[len(a)-1] * b[len(b)-1])

	}
}

func NewCF1631A(r *bufio.Reader) *CF1631A {
	return &CF1631A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1631A(bufio.NewReader(os.Stdin)).Run()
}
