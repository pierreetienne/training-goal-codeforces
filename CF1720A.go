package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1720A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1720A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1720A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1720A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1720A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1720A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1720A
Date: 8/27/2022
User: wotan
URL: https://codeforces.com/contest/1720/problem/A
Problem: CF1720A
**/
func (in *CF1720A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, b, c, d := in.NextInt64(), in.NextInt64(), in.NextInt64(), in.NextInt64()
		x := a * d
		y := b * c
		if x == y {
			fmt.Println(0)
		} else if y != 0 && x%y == 0 || x != 0 && y%x == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}
	}
}

func NewCF1720A(r *bufio.Reader) *CF1720A {
	return &CF1720A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1720A(bufio.NewReader(os.Stdin)).Run()
}
