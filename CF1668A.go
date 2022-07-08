package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1668A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1668A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1668A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1668A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1668A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1668A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1668A
Date: 6/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1668/A
Problem: CF1668A
**/
func (in *CF1668A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		x, y := in.NextInt(), in.NextInt()

		a := 0
		b := 0
		if x < y {
			a = x
			b = y
		} else {
			a = y
			b = x
		}

		if a == 1 && b >= 3 {
			fmt.Println(-1)
		} else {
			fmt.Println(2*b - 2 - (b+a)%2)
		}
	}
}

func NewCF1668A(r *bufio.Reader) *CF1668A {
	return &CF1668A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1668A(bufio.NewReader(os.Stdin)).Run()
}
