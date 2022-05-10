package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1674A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1674A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1674A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1674A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1674A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1674A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1674A
Date: 6/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1674/A
Problem: CF1674A
**/
func (in *CF1674A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		x, y := in.NextInt(), in.NextInt()

		p, q := 0, 0
		if y%x == 0 {
			p = 1
			q = y / x
		}
		fmt.Println(p, q)
	}
}

func NewCF1674A(r *bufio.Reader) *CF1674A {
	return &CF1674A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1674A(bufio.NewReader(os.Stdin)).Run()
}
