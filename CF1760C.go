package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1760C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1760C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1760C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1760C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1760C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1760C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1760C
Date: 11/21/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1760/C
Problem: CF1760C
**/
func (in *CF1760C) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()

		arr := make([]int, n)
		arr2 := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			arr2[i] = arr[i]
		}
		sort.Ints(arr2)
		for i := 0; i < n; i++ {
			if arr[i] == arr2[n-1] {
				fmt.Print(arr[i]-arr2[n-2], " ")
			} else {
				fmt.Print(arr[i]-arr2[n-1], " ")
			}
		}
		fmt.Println()
	}
}

func NewCF1760C(r *bufio.Reader) *CF1760C {
	return &CF1760C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1760C(bufio.NewReader(os.Stdin)).Run()
}
