package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1393A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1393A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1393A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1393A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1393A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1393A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1393A
Date: 20/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1393/A
Problem: CF1393A
**/
func (in *CF1393A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		if n == 1 || n == 2 {
			fmt.Println(n)
		} else {
			sol := int(math.Ceil(float64(n+1) / 2.0))
			fmt.Println(sol)
		}

	}
}

func NewCF1393A(r *bufio.Reader) *CF1393A {
	return &CF1393A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1393A(bufio.NewReader(os.Stdin)).Run()
}
