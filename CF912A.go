package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF912A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF912A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF912A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF912A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF912A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF912A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF912A
Date: 31/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/912/A
Problem: CF912A
**/
func (in *CF912A) Run() {
	A, B := in.NextInt64(), in.NextInt64()
	x, y, z := in.NextInt64(), in.NextInt64(), in.NextInt64()

	A -= x * 2
	B = B - y
	A = A - y
	B -= 3 * z

	ans := int64(0)
	if A < 0 {
		ans = -A
	}
	if B < 0 {
		ans += -B
	}
	fmt.Println(ans)
}

func NewCF912A(r *bufio.Reader) *CF912A {
	return &CF912A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF912A(bufio.NewReader(os.Stdin)).Run()
}
