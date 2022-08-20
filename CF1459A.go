package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1459A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1459A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1459A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1459A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1459A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1459A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1459A
Date: 8/16/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1459/A
Problem: CF1459A
**/
func (in *CF1459A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		r := in.NextString()
		b := in.NextString()

		x, y := 0, 0
		for i := 0; i < n; i++ {
			if b[i] > r[i] {
				x++
			}
			if r[i] > b[i] {
				y++
			}
		}
		if y == x {
			fmt.Println("EQUAL")
		} else if x > y {
			fmt.Println("BLUE")
		} else {
			fmt.Println("RED")
		}
	}
}

func NewCF1459A(r *bufio.Reader) *CF1459A {
	return &CF1459A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1459A(bufio.NewReader(os.Stdin)).Run()
}
