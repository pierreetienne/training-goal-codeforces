package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1701A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1701A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1701A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1701A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1701A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1701A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1701A
Date: 17/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1701/A
Problem: CF1701A
**/
func (in *CF1701A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a := in.NextString()+" "+in.NextString()
		b := in.NextString()+" "+in.NextString()

		ans := 0
		if a == "0 0" && b == "0 0" {
			ans = 0
		}else if a == "1 1" && b == "1 1" {
			ans = 2
		}else {
			ans  = 1
		}
		fmt.Println(ans)
	}
}

func NewCF1701A(r *bufio.Reader) *CF1701A {
	return &CF1701A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1701A(bufio.NewReader(os.Stdin)).Run()
}
