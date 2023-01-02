package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1731A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1731A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1731A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1731A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1731A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1731A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1731A
Date: 12/27/2022
User: wotan
URL: https://codeforces.com/contests/1731A
Problem: CF1731A
**/
func (in *CF1731A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		sol := int64(1)
		for i := 0; i < n; i++ {
			sol *= in.NextInt64()
		}
		fmt.Println((sol + int64(n-1)) * 2022)
	}
}

func NewCF1731A(r *bufio.Reader) *CF1731A {
	return &CF1731A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1731A(bufio.NewReader(os.Stdin)).Run()
}
