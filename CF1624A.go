package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1624A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1624A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1624A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1624A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1624A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1624A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1626A
Date: 12/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1624/A
Problem: CF1624A
**/
func (in *CF1624A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		max := 0
		min := 0
		for i := 0; i < n; i++ {
			num := in.NextInt()
			if i == 0 {
				min = num
				max = num
			} else {
				if num > max {
					max = num
				}

				if num < min {
					min = num
				}
			}
		}

		fmt.Println(max - min)

	}
}

func NewCF1624A(r *bufio.Reader) *CF1626A {
	return &CF1626A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1626A(bufio.NewReader(os.Stdin)).Run()
}
