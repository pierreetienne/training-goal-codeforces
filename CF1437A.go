package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1437A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1437A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1437A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1437A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1437A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1437A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1437A
Date: 14/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1437/A
Problem: CF1437A
**/
func (in *CF1437A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		l, r := in.NextInt(), in.NextInt()

		a:= r+1
		if float64(l % a) >= float64(a) / 2.0 && float64(r % a) >= float64(a) / 2.0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func NewCF1437A(r *bufio.Reader) *CF1437A {
	return &CF1437A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1437A(bufio.NewReader(os.Stdin)).Run()
}
