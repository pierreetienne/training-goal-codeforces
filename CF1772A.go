package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1772A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1772A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1772A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1772A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1772A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1772A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1772A
Date: 12/19/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1772/A
Problem: CF1772A
**/
func (in *CF1772A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		fmt.Println((int)(str[0]-'0') + (int)(str[2]-'0'))
	}
}

func NewCF1772A(r *bufio.Reader) *CF1772A {
	return &CF1772A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1772A(bufio.NewReader(os.Stdin)).Run()
}
