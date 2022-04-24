package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1669B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1669B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1669B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1669B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1669B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1669B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1669B
Date: 22/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1669/B
Problem: CF1669B
**/
func (in *CF1669B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		sort.Ints(arr)
		ans := -1
		for i := 2; i < n; i++ {
			if arr[i] == arr[i-1] && arr[i] == arr[i-2] {
				ans = arr[i]
				break
			}
		}
		fmt.Println(ans)
	}
}

func NewCF1669B(r *bufio.Reader) *CF1669B {
	return &CF1669B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1669B(bufio.NewReader(os.Stdin)).Run()
}
