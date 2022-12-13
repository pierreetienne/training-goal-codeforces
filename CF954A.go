package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF954A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF954A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF954A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF954A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF954A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF954A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF954A
Date: 12/11/2022
User: wotan
URL: https://codeforces.com/problemset/problem/954/A
Problem: CF954A
**/
func (in *CF954A) Run() {
	n := in.NextInt()
	str := in.NextString()

	c := 0
	for i := 1; i < n; i++ {
		if str[i-1] == 'R' && str[i] == 'U' {
			c++
			i++
		} else if str[i-1] == 'U' && str[i] == 'R' {
			c++
			i++
		}
	}

	fmt.Println(n - c)
}

func NewCF954A(r *bufio.Reader) *CF954A {
	return &CF954A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF954A(bufio.NewReader(os.Stdin)).Run()
}
