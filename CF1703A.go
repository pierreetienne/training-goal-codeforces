package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1703A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1703A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1703A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1703A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1703A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1703A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1703A
Date: 13/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1703/A
Problem: CF1703A
**/
func (in *CF1703A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		if strings.ToLower(in.NextString()) == "yes" {
			fmt.Println("YES")
		}else {
			fmt.Println("NO")
		}
	}
}

func NewCF1703A(r *bufio.Reader) *CF1703A {
	return &CF1703A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1703A(bufio.NewReader(os.Stdin)).Run()
}
