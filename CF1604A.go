package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1604A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1604A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1604A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1604A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1604A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1604A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1604A
Date: 14/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1604/A
Problem: CF1604A
**/
func (in *CF1604A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		index := 1
		ans := 0
		for i := 0; i < n; i++ {
			num := in.NextInt()
			if index < num {
				ans += num - index
				index += num - index
			}
			index++
		}
		fmt.Println(ans)
	}
}

func NewCF1604A(r *bufio.Reader) *CF1604A {
	return &CF1604A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1604A(bufio.NewReader(os.Stdin)).Run()
}
