package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF157A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF157A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF157A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF157A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF157A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF157A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF157A
Date: 1/17/2023
User: wotan
URL: https://codeforces.com/problemset/problem/157/A
Problem: CF157A
**/
func (in *CF157A) Run() {
	n := in.NextInt()
	cols := make([]int, n)
	rows := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := in.NextInt()
			rows[i] += val
			cols[j] += val
		}
	}

	ans := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if cols[j] > rows[i] {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func NewCF157A(r *bufio.Reader) *CF157A {
	return &CF157A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF157A(bufio.NewReader(os.Stdin)).Run()
}
