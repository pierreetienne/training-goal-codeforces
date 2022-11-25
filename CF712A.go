package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF712A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF712A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF712A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF712A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF712A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF712A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots timelimit for output use string.builder
Run solve the problem CF712A
Date: 11/23/2022
User: wotan
URL: https://codeforces.com/problemset/problem/712/A
Problem: CF712A
**/
func (in *CF712A) Run() {
	n := in.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
	}
	s := make([]int, n)
	var sol strings.Builder
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			s[i] = arr[i]
		} else {
			s[i] = arr[i] + arr[i+1]
		}

	}
	for i := 0; i < n; i++ {
		sol.WriteString(fmt.Sprintf("%v ", s[i]))
	}
	fmt.Println(sol.String())
}

func NewCF712A(r *bufio.Reader) *CF712A {
	return &CF712A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF712A(bufio.NewReader(os.Stdin)).Run()
}
