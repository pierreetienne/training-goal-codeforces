package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF104A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF104A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF104A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF104A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF104A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF104A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF104A
Date: 11/30/2022
User: wotan
URL: https://codeforces.com/problemset/problem/104/A
Problem: CF104A
**/
func (in *CF104A) Run() {
	n := in.NextInt()
	if n <= 10 {
		fmt.Println(0)
	} else if n <= 19 {
		fmt.Println(4)
	} else if n == 20 {
		fmt.Println(15)
	} else if n == 21 {
		fmt.Println(4)
	} else {
		fmt.Println(0)
	}
}

func NewCF104A(r *bufio.Reader) *CF104A {
	return &CF104A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF104A(bufio.NewReader(os.Stdin)).Run()
}
