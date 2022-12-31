package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF664A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF664A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF664A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF664A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF664A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF664A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF664A
Date: 11/4/2022
User: wotan
URL: https://codeforces.com/problemset/problem/664/A
Problem: CF664A
**/
func (in *CF664A) Run() {
	a, b := in.NextString(), in.NextString()
	if a == b {
		fmt.Println(a)
	} else {
		fmt.Println(1)
	}
}

func NewCF664A(r *bufio.Reader) *CF664A {
	return &CF664A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF664A(bufio.NewReader(os.Stdin)).Run()
}
