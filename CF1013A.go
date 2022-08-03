package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1013A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1013A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1013A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1013A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1013A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1013A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1013A
Date: 1/08/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1013/A
Problem: CF1013A
**/
func (in *CF1013A) Run() {
	n := in.NextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		val := in.NextInt()
		a[i] = val
	}

	ws := true
	minus := 0
	add := 0
	for i := 0; i < n; i++ {
		val := in.NextInt()
		if val > a[i] {
			add += val - a[i]
		} else if val < a[i] {
			minus += a[i] - val
		}
	}
	ws = add <= minus
	if ws {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func NewCF1013A(r *bufio.Reader) *CF1013A {
	return &CF1013A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1013A(bufio.NewReader(os.Stdin)).Run()
}
