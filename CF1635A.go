package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1635A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1635A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1635A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1635A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1635A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1635A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1635A
Date: 17/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1635/A
Problem: CF1635A
**/
func (in *CF1635A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		ans := 0
		for i := 0; i < n; i++ {
			ans |= in.NextInt()
		}
		fmt.Println(ans)
	}
}

func NewCF1635A(r *bufio.Reader) *CF1635A {
	return &CF1635A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1635A(bufio.NewReader(os.Stdin)).Run()
}
