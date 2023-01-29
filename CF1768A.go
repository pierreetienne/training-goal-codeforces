package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1768A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1768A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1768A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1768A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1768A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1768A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1768A
Date: 1/26/2023
User: wotan
URL: https://codeforces.com/problemset/problem/1768/A
Problem: CF1768A
**/
func (in *CF1768A) Run() {
	var str strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		str.WriteString(fmt.Sprintf("%d\n", in.NextInt()-1))
	}
	fmt.Print(str.String())
}

func NewCF1768A(r *bufio.Reader) *CF1768A {
	return &CF1768A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1768A(bufio.NewReader(os.Stdin)).Run()
}
