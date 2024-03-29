package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1761A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1761A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1761A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1761A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1761A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1761A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1761A
Date: 11/23/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1761/A
Problem: CF1761A
**/
func (in *CF1761A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, a, b := in.NextInt(), in.NextInt(), in.NextInt()
		if n-a-b >= 2 || (n == a && a == b) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func NewCF1761A(r *bufio.Reader) *CF1761A {
	return &CF1761A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1761A(bufio.NewReader(os.Stdin)).Run()
}
