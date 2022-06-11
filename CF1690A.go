package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1690A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1690A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1690A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1690A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1690A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1690A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1690A
Date: 8/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1690/A
Problem: CF1690A
**/
func (in *CF1690A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		a, b, c := 2, n-3, 1
		for i := b; i >= 0; i-- {
			if i-1 > 0 && n-i-(i-1) > 0 && n-i-(i-1) < i-1 {
				a = i - 1
				b = i
				c = n - i - (i - 1)
			}
		}
		fmt.Println(a, b, c)
	}
}

func NewCF1690A(r *bufio.Reader) *CF1690A {
	return &CF1690A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1690A(bufio.NewReader(os.Stdin)).Run()
}
