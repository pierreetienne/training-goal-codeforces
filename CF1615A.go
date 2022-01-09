package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1615A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1615A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1615A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1615A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1615A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1615A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1615A
Date: 8/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1615/A
Problem: CF1615A
**/
func (in *CF1615A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		sum := 0
		for i := 0; i < n; i++ {
			num := in.NextInt()
			sum += num
		}

		ans := 0
		if sum%n != 0 {
			ans = 1
		}
		fmt.Println(ans)
	}
}

func NewCF1615A(r *bufio.Reader) *CF1615A {
	return &CF1615A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1615A(bufio.NewReader(os.Stdin)).Run()
}
