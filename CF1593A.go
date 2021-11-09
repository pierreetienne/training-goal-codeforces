package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1593A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1593A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1593A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1593A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1593A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1593A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1593A
Date: 8/11/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1593/A
Problem: CF1593A
**/
func (in *CF1593A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		a, b, c := in.NextInt64(), in.NextInt64(), in.NextInt64()

		if a > b && a > c {
			fmt.Print(0, " ")
		} else {
			max := max(b,c)+1
			fmt.Print(max-a, " ")
		}

		if b > a && b > c {
			fmt.Print(0, " ")
		} else {
			max := max(a,c)+1
			fmt.Print(max-b, " ")
		}

		if c > a && c > b {
			fmt.Print(0, " ")
		} else {
			max := max(a,b)+1
			fmt.Print(max-c, " ")
		}
		fmt.Println()
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func NewCF1593A(r *bufio.Reader) *CF1593A {
	return &CF1593A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1593A(bufio.NewReader(os.Stdin)).Run()
}
