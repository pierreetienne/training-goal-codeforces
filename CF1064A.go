package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1064A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1064A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1064A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1064A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1064A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1064A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1064A
Date: 14/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1064/A
Problem: CF1064A
**/
func (in *CF1064A) Run() {
	a, b, c := in.NextInt(), in.NextInt(), in.NextInt()

	if a+b > c && b+c > a && a+c > b {
		fmt.Println(0)
	} else {
		ans := 0
		if a+b <= c {
			ans += c - (a + b) + 1
		} else if a+c <= b {
			ans += b - (a + c) + 1
		} else {
			ans += a - (b + c) + 1
		}
		fmt.Println(ans)
	}

}

func NewCF1064A(r *bufio.Reader) *CF1064A {
	return &CF1064A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1064A(bufio.NewReader(os.Stdin)).Run()
}
