package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1728A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1728A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1728A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1728A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1728A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1728A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1728A
Date: 9/9/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1728/A
Problem: CF1728A
**/
func (in *CF1728A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()

		max := -1
		color := 0
		for i := 0; i < n; i++ {
			val := in.NextInt()
			if val > max {
				max = val
				color = i + 1
			}
		}

		fmt.Println(color)
	}
}

func NewCF1728A(r *bufio.Reader) *CF1728A {
	return &CF1728A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1728A(bufio.NewReader(os.Stdin)).Run()
}
