package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1660B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1660B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1660B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1660B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1660B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1660B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1660B
Date: 3/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1660/B
Problem: CF1660B
**/
func (in *CF1660B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		sum := int64(0)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			sum += int64(arr[i])
		}

		ans := false
		if n == 1 {
			if sum == 1 {
				ans = true
			}
		} else {
			sort.Ints(arr)
			ws := true
			if arr[n-1]-arr[n-2] > 1 {
				ws = false
			}

			ans = ws
		}

		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1660B(r *bufio.Reader) *CF1660B {
	return &CF1660B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1660B(bufio.NewReader(os.Stdin)).Run()
}
