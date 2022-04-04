package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1660A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1660A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1660A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1660A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1660A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1660A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1660A
Date: 2/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1660/A
Problem: CF1660A
**/
func (in *CF1660A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, b := in.NextInt(), in.NextInt()
		if a == 0 {
			fmt.Println(1)
		} else {
			fmt.Println((b * 2) + a + 1)
		}
	}
}

func NewCF1660A(r *bufio.Reader) *CF1660A {
	return &CF1660A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1660A(bufio.NewReader(os.Stdin)).Run()
}
