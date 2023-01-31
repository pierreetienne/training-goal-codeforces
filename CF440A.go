package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF440A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF440A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF440A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF440A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF440A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF440A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF440A
Date: 1/27/2023
User: wotan
URL: https://codeforces.com/problemset/problem/440/A
Problem: CF440A
**/
func (in *CF440A) Run() {
	n := in.NextInt()

	f := (n * (n + 1)) / 2
	for i := 0; i < n-1; i++ {
		f -= in.NextInt()
	}
	fmt.Println(f)
}

func NewCF440A(r *bufio.Reader) *CF440A {
	return &CF440A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF440A(bufio.NewReader(os.Stdin)).Run()
}
