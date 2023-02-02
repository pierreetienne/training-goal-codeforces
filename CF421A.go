package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF421A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF421A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF421A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF421A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF421A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF421A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF421A
Date: 1/30/2023
User: wotan
URL: https://codeforces.com/problemset/problem/421/A
Problem: CF421A
**/
func (in *CF421A) Run() {
	n, a, b := in.NextInt(), in.NextInt(), in.NextInt()
	arr := make([]string, n)

	for i := 0; i < a; i++ {
		arr[in.NextInt()-1] = "1"
	}
	for i := 0; i < b; i++ {
		arr[in.NextInt()-1] = "2"
	}

	var str strings.Builder
	for i := 0; i < n; i++ {
		str.WriteString(arr[i] + " ")
	}
	fmt.Println(str.String())
}

func NewCF421A(r *bufio.Reader) *CF421A {
	return &CF421A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF421A(bufio.NewReader(os.Stdin)).Run()
}
