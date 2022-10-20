package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1743A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1743A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1743A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1743A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1743A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1743A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1743A
Date: 10/18/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1743/A
Problem: CF1743A
**/
func (in *CF1743A) Run() {

	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		for i := 0; i < n; i++ {
			in.NextInt()
		}

		sol := fib(10-n-1) * 6
		fmt.Println(sol)

	}

}

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return n + fib(n-1)
}

func NewCF1743A(r *bufio.Reader) *CF1743A {
	return &CF1743A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1743A(bufio.NewReader(os.Stdin)).Run()
}
