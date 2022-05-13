package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1676A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1676A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1676A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1676A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1676A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1676A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1676A
Date: 11/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1676/A
Problem: CF1676A
**/
func (in *CF1676A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		s := 0
		a := 0
		for i := 0; i < len(str); i++ {
			if i < 3 {
				s += int(str[i] - '0')
			}
			a += int(str[i] - '0')
		}

		if s == (a - s) {
			fmt.Println("YES")
		}else {
			fmt.Println("NO")
		}
	}
}

func NewCF1676A(r *bufio.Reader) *CF1676A {
	return &CF1676A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1676A(bufio.NewReader(os.Stdin)).Run()
}
