package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1712A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1712A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1712A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1712A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1712A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1712A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1712A
Date: 8/20/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1712/B
Problem: CF1712A
**/
func (in *CF1712A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()

		if n == 1 {
			fmt.Println(1)
		} else if n >= 2 {
			if n%2 == 0 {
				for i := 1; i <= n; i += 2 {
					fmt.Print(i+1, i, " ")
				}
			} else {
				fmt.Print(1, " ")
				for i := 3; i <= n; i += 2 {
					fmt.Print(i, i-1, " ")
				}
			}
			fmt.Println()
		}
	}
}

func NewCF1712A(r *bufio.Reader) *CF1712A {
	return &CF1712A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1712A(bufio.NewReader(os.Stdin)).Run()
}
