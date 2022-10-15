package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1735A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1735A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1735A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1735A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1735A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1735A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1735A
Date: 10/6/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1735/A
Problem: CF1735A
**/
func (in *CF1735A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()

		fmt.Println((n - 3/3) - 1)
	}
}

func NewCF1735A(r *bufio.Reader) *CF1735A {
	return &CF1735A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1735A(bufio.NewReader(os.Stdin)).Run()
}
