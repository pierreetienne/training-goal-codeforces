package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1598A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1598A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1598A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1598A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1598A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1598A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1598A
Date: 16/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1598/A
Problem: CF1598A
**/
func (in *CF1598A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		a := in.NextString()
		b := in.NextString()
		ans := true
		for i := 0; i < n; i++ {
			if a[i] == '1' && b[i] == '1' {
				ans = false
				break
			}
		}
		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1598A(r *bufio.Reader) *CF1598A {
	return &CF1598A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1598A(bufio.NewReader(os.Stdin)).Run()
}
