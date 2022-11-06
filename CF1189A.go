package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1189A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1189A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1189A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1189A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1189A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1189A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1189A
Date: 11/3/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1189/A
Problem: CF1189A
**/
func (in *CF1189A) Run() {
	n := in.NextInt()
	s := in.NextString()
	ones := make([]int, n)
	zeros := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			zeros[i] = 1
		} else {
			ones[i] = 1
		}
		if i > 0 {
			zeros[i] += zeros[i-1]
			ones[i] += ones[i-1]
		}
	}

	if ones[n-1] != zeros[n-1] {
		fmt.Println(1)
		fmt.Println(s)
	} else {
		fmt.Println(2)
		fmt.Println(s[0:n-1], s[n-1:])
	}

}

func NewCF1189A(r *bufio.Reader) *CF1189A {
	return &CF1189A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1189A(bufio.NewReader(os.Stdin)).Run()
}
