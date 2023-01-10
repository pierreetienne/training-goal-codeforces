package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF447A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF447A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF447A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF447A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF447A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF447A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF447A
Date: 1/2/2023
User: wotan
URL: https://codeforces.com/problemset/problem/447/A
Problem: CF447A
**/
func (in *CF447A) Run() {
	p, n := in.NextInt(), in.NextInt()
	arr := make([]bool, p)
	ans := -1
	for i := 0; i < n; i++ {
		x := in.NextInt()
		val := x % p
		if !arr[val] {
			arr[val] = true
		} else if ans == -1 {
			ans = i + 1
		}
	}
	fmt.Println(ans)
}

func NewCF447A(r *bufio.Reader) *CF447A {
	return &CF447A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF447A(bufio.NewReader(os.Stdin)).Run()
}
