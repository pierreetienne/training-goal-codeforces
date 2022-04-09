package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF172A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF172A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF172A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF172A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF172A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF172A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF172A
Date: 7/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/172/A
Problem: CF172A
**/
func (in *CF172A) Run() {
	n := in.NextInt()
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextString()
	}

	sol := 0
	for i := 0; i < len(arr[0]); i++ {
		ok := true
		for j := 1; j < n; j++ {
			if arr[j][i] != arr[j-1][i] {
				ok = false
			}
		}
		if ok {
			sol++
		} else {
			break
		}
	}

	fmt.Println(sol)
}

func NewCF172A(r *bufio.Reader) *CF172A {
	return &CF172A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF172A(bufio.NewReader(os.Stdin)).Run()
}
