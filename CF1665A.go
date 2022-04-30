package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1665A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1665A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1665A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1665A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1665A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1665A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1665A
Date: 28/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1665/A
Problem: CF1665A
**/
func (in *CF1665A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		fmt.Println(1, n-3, 1, 1)

	}
}

func NewCF1665A(r *bufio.Reader) *CF1665A {
	return &CF1665A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1665A(bufio.NewReader(os.Stdin)).Run()
}
