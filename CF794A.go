package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF794A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF794A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF794A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF794A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF794A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF794A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF794A
Date: 1/16/2023
User: wotan
URL: https://codeforces.com/problemset/problem/794/A
Problem: CF794A
**/
func (in *CF794A) Run() {
	_, b, c := in.NextInt(), in.NextInt(), in.NextInt()
	n := in.NextInt()
	sol := int64(0)
	for i := 0; i < n; i++ {
		val := in.NextInt()
		if b < val && val < c {
			sol++
		}
	}

	fmt.Println(sol)

}

func NewCF794A(r *bufio.Reader) *CF794A {
	return &CF794A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF794A(bufio.NewReader(os.Stdin)).Run()
}
