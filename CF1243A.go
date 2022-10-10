package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1243A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1243A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1243A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1243A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1243A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1243A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
TO SEND
Run solve the problem CF1243A
Date: 7/10/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1243/A
Problem: CF1243A
**/
func (in *CF1243A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		sort.Ints(arr)
		max := 0
		for i := 0; i < n; i++ {
			if arr[i] >= (n - i) {
				max = n - i
				break
			}
		}
		fmt.Println(max)
	}
}

func NewCF1243A(r *bufio.Reader) *CF1243A {
	return &CF1243A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1243A(bufio.NewReader(os.Stdin)).Run()
}
