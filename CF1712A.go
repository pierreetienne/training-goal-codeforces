package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1712A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1712A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1712A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1712A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1712A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1712A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1712A
Date: 18/08/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1712/A
Problem: CF1712A
**/
func (in *CF1712A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()
		p := make([]int, n)
		mapa := make(map[int]int)
		for i := 0; i < n; i++ {
			p[i] = in.NextInt()
			if i < k {
				mapa[p[i]]++
			}
		}
		sort.Ints(p)
		ans := 0
		for i := 0; i < k; i++ {
			if v, e := mapa[p[i]]; e {
				mapa[p[i]] = v - 1
			} else {
				ans++
			}
		}
		fmt.Println(ans)
	}
}

func NewCF1712A(r *bufio.Reader) *CF1712A {
	return &CF1712A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1712A(bufio.NewReader(os.Stdin)).Run()
}
