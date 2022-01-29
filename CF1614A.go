package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1614A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1614A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1614A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1614A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1614A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1614A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
// gopherbots contest
Run solve the problem CF1614A
Date: 27/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1614/A
Problem: CF1614A
**/
func (in *CF1614A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, l, r, k := in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		sort.Ints(arr)
		i, j := 0, n-1
		for {
			move := false
			if i < n && arr[i] < l {
				i++
				move = true
			}
			if j >= 0 && arr[j] > r {
				j--
				move = true
			}
			if !move {
				break
			}
		}

		ans := 0
		for p := i; p <= j; p++ {
			if k > 0 && k-arr[p] >= 0 {
				k -= arr[p]
				ans++
			} else {
				break
			}
		}
		fmt.Println(ans)
	}
}

func NewCF1614A(r *bufio.Reader) *CF1614A {
	return &CF1614A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1614A(bufio.NewReader(os.Stdin)).Run()
}
