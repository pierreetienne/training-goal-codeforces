package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1730A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1730A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1730A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1730A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1730A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1730A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1730A
Date: 1/10/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1730/A
Problem: CF1730A
**/
func (in *CF1730A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, c := in.NextInt(), in.NextInt()
		arr := make([]int, 101)
		for i := 0; i < n; i++ {
			arr[in.NextInt()]++
		}
		cost := 0
		for i := 0; i < len(arr); i++ {
			if arr[i] > 0 {
				if arr[i] > c {
					cost += c
				} else {
					cost += arr[i]
				}
			}
		}
		fmt.Println(cost)

	}
}

func NewCF1730A(r *bufio.Reader) *CF1730A {
	return &CF1730A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1730A(bufio.NewReader(os.Stdin)).Run()
}
