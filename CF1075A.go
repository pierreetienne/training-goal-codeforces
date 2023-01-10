package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1075A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1075A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1075A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1075A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1075A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1075A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1075A
Date: 1/6/2023
User: wotan
URL: https://codeforces.com/problemset/problem/1075/A
Problem: CF1075A
**/
func (in *CF1075A) Run() {
	n := in.NextInt64()
	x, y := in.NextInt64(), in.NextInt64()
	if x+y-2 > 2*n-x-y {
		fmt.Println("Black")
	} else {
		fmt.Println("White")
	}
}

func NewCF1075A(r *bufio.Reader) *CF1075A {
	return &CF1075A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1075A(bufio.NewReader(os.Stdin)).Run()
}
