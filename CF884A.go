package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF884A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF884A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF884A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF884A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF884A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF884A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF884A
Date: 8/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/884/A
Problem: CF884A
**/
func (in *CF884A) Run() {
	n, t := in.NextInt(), in.NextInt()
	ans := 0
	for i := 0; i < n; i++ {
		val := in.NextInt()
		if t > 0 {
			t = t - (86400 - val)
			ans++
		}
	}
	fmt.Println(ans)
}

func NewCF884A(r *bufio.Reader) *CF884A {
	return &CF884A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF884A(bufio.NewReader(os.Stdin)).Run()
}
