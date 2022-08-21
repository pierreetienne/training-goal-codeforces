package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1719A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1719A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1719A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1719A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1719A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1719A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1719A
Date: 8/19/2022
User: wotan
URL: https://codeforces.com/contest/1719/problem/A
Problem: CF1719A
**/
func (in *CF1719A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt64(), in.NextInt64()
		if (n+m)%2 == 0 {
			fmt.Println("Tonya")
		} else {
			fmt.Println("Burenka")
		}
	}
}

func NewCF1719A(r *bufio.Reader) *CF1719A {
	return &CF1719A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1719A(bufio.NewReader(os.Stdin)).Run()
}
