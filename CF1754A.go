package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1754A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1754A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1754A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1754A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1754A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1754A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1754A
Date: 10/24/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1754/A
Problem: CF1754A
**/
func (in *CF1754A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		q := 0

		for i := 0; i < n; i++ {
			if str[i] == 'Q' {
				q++
			} else {
				q--
				if q < 0 {
					q = 0
				}
			}
		}
		if q > 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}

func NewCF1754A(r *bufio.Reader) *CF1754A {
	return &CF1754A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1754A(bufio.NewReader(os.Stdin)).Run()
}
