package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1758A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1758A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1758A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1758A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1758A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1758A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1758A
Date: 11/25/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1758/A
Problem: CF1758A
**/
func (in *CF1758A) Run() {
	c := int64(1)
	for i := 1; i <= 18; i++ {
		c *= int64(2) % 1000000007
		fmt.Println(i, c)
	}
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		out := ""
		for i := 0; i < len(str); i++ {
			out = string(str[i]) + out + string(str[i])
		}
		fmt.Println(out)
	}
}

func NewCF1758A(r *bufio.Reader) *CF1758A {
	return &CF1758A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1758A(bufio.NewReader(os.Stdin)).Run()
}
