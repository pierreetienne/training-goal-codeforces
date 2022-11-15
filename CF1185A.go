package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1185A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1185A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1185A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1185A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1185A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1185A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1185A
Date: 6/11/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1185/A
Problem: CF1185A
**/
func (in *CF1185A) Run() {
	arr := []int{in.NextInt(), in.NextInt(), in.NextInt()}
	sort.Ints(arr)
	d := in.NextInt()
	ans := 0
	if arr[1]-arr[0] < d {
		ans += d - (arr[1] - arr[0])
	}
	if arr[2]-arr[1] < d {
		ans += d - (arr[2] - arr[1])
	}
	fmt.Println(ans)

}

func NewCF1185A(r *bufio.Reader) *CF1185A {
	return &CF1185A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1185A(bufio.NewReader(os.Stdin)).Run()
}
