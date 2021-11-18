package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1451A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1451A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1451A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1451A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1451A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1451A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1451A
Date: 17/11/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1451/A
Problem: CF1451A
**/
func (in *CF1451A) Run() {

	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		if n == 1 {
			fmt.Println(0)
		} else if n == 2 {
			fmt.Println(1)
		} else if n == 3 {
			fmt.Println(2)
		} else {
			if n%2 == 0 {
				fmt.Println(2)
			} else {
				fmt.Println(3)
			}
		}

	}
}

func NewCF1451A(r *bufio.Reader) *CF1451A {
	return &CF1451A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1451A(bufio.NewReader(os.Stdin)).Run()
}
