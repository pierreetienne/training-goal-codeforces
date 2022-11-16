package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1748A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1748A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1748A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1748A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1748A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1748A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1748A
Date: 11/14/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1748/A
Problem: CF1748A
**/
func (in *CF1748A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		fmt.Println(int(math.Ceil(float64(n) / 2.)))
	}
}

func NewCF1748A(r *bufio.Reader) *CF1748A {
	return &CF1748A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1748A(bufio.NewReader(os.Stdin)).Run()
}
