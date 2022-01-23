package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1499A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1499A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1499A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1499A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1499A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1499A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1499A
Date: 22/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1499/A
Problem: CF1499A
**/
func (in *CF1499A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k1, k2 := in.NextInt(), in.NextInt(), in.NextInt()
		w, b := in.NextInt(), in.NextInt()
		k1b, k2b := n-k1, n-k2
		ans := "YES"
		if k1+k2 >= 2*w && k1b+k2b >= 2*b || (w == 0 && b == 0) {

		} else {
			ans = "NO"
		}
		fmt.Println(ans)
	}
}

func NewCF1499A(r *bufio.Reader) *CF1499A {
	return &CF1499A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1499A(bufio.NewReader(os.Stdin)).Run()
}
