package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF32A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF32A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF32A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF32A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF32A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF32A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF32A
Date: 12/13/2022
User: wotan
URL: https://codeforces.com/problemset/problem/32/A
Problem: CF32A
**/
func (in *CF32A) Run() {
	n, d := in.NextInt(), in.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
	}

	sort.Ints(arr)

	cont := 0
	for i := n - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if arr[i]-arr[j] <= d {
				cont += 2
			} else {
				break
			}
		}
	}

	fmt.Println(cont)
}

func NewCF32A(r *bufio.Reader) *CF32A {
	return &CF32A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF32A(bufio.NewReader(os.Stdin)).Run()
}
