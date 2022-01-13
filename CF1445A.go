package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1445A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1445A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		//fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1445A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1445A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1445A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1445A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1445A
Date: 13/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1445/A
Problem: CF1445A
**/
func (in *CF1445A) Run() {

	for t := in.NextInt(); t > 0; t-- {
		n, x := in.NextInt(), in.NextInt()
		a := make([]int, n)
		b := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = in.NextInt()
		}
		for i := 0; i < n; i++ {
			b[i] = in.NextInt()
		}
		in.GetLine()
		sort.Ints(a)
		sort.Ints(b)
		ws := true
		for i, j := 0, n-1; i < n; i++ {
			if a[i]+b[j] > x {
				ws = false
				break
			}
			j--
		}

		if ws {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func NewCF1445A(r *bufio.Reader) *CF1445A {
	return &CF1445A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1445A(bufio.NewReader(os.Stdin)).Run()
}
