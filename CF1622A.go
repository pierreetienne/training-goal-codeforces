package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1622A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1622A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1622A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1622A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1622A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1622A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1622A
Date: 3/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1622/A
Problem: CF1622A
**/
func (in *CF1622A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		arr := []int{in.NextInt(), in.NextInt(), in.NextInt()}
		sort.Ints(arr)
		ws := false
		if (arr[0] == arr[1] && arr[2]%2 == 0) || (arr[1] == arr[2] && arr[0]%2 == 0) {
			ws = true
		} else if arr[2]-arr[1]-arr[0] == 0 {
			ws = true
		}

		if ws {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1622A(r *bufio.Reader) *CF1622A {
	return &CF1622A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1622A(bufio.NewReader(os.Stdin)).Run()
}
