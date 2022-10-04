package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF888A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF888A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF888A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF888A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF888A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF888A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF888A
Date: 6/08/22
User: pepradere
URL: https://codeforces.com/problemset/problem/888/A
Problem: CF888A
**/
func (in *CF888A) Run() {
	n := in.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			ans++
		} else if arr[i-1] > arr[i] && arr[i] < arr[i+1] {
			ans++
		}
	}

	fmt.Println(ans)
}

func NewCF888A(r *bufio.Reader) *CF888A {
	return &CF888A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF888A(bufio.NewReader(os.Stdin)).Run()
}
