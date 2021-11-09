package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1480A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1480A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1480A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1480A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1480A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1480A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1480A
Date: 8/11/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1480/A
Problem: CF1480A
**/
func (in *CF1480A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		str := in.NextString()
		sol := ""
		for i := 0; i < len(str); i++ {
			if i%2 == 0 {
				if str[i] != 'a' {
					sol += "a"
				} else {
					sol += "b"
				}
			} else {
				if str[i] != 'z' {
					sol += "z"
				} else {
					sol += "y"
				}
			}
		}
		fmt.Println(sol)
	}
}

func NewCF1480A(r *bufio.Reader) *CF1480A {
	return &CF1480A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1480A(bufio.NewReader(os.Stdin)).Run()
}
