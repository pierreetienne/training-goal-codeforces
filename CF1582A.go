package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1582A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1582A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1582A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1582A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1582A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1582A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1582A
Date: 24/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1582/A
Problem: CF1582A
**/
func (in *CF1582A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, _, c := in.NextInt(), in.NextInt(), in.NextInt()
		fmt.Println((a + c) % 2)
	}
}

func NewCF1582A(r *bufio.Reader) *CF1582A {
	return &CF1582A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1582A(bufio.NewReader(os.Stdin)).Run()
}
