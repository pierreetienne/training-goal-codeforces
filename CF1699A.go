package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1699A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1699A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1699A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1699A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1699A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1699A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1699A
Date: 10/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1699/A
Problem: CF1699A
**/
func (in *CF1699A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt64()
		if n%2 == 0 {
			fmt.Println( n/2, n/2, 0)
		} else {
			fmt.Println(-1)
		}

	}
}

func NewCF1699A(r *bufio.Reader) *CF1699A {
	return &CF1699A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1699A(bufio.NewReader(os.Stdin)).Run()
}
