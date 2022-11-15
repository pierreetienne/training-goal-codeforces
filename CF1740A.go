package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1740A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1740A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1740A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1740A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1740A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1740A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1740A
Date: 7/11/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1740/A
Problem: CF1740A
**/
func (in *CF1740A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		in.NextInt()
		fmt.Println(7)
	}
}

func NewCF1740A(r *bufio.Reader) *CF1740A {
	return &CF1740A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1740A(bufio.NewReader(os.Stdin)).Run()
}
