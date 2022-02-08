package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1632A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1632A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1632A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1632A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1632A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1632A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1632A
Date: 6/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1632/A
Problem: CF1632A
**/
func (in *CF1632A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		zeros := 0
		for i := 0; i < n; i++ {
			if str[i] == '0' {
				zeros++
			}
		}
		ones := n - zeros
		if (ones == 1 && zeros == 1) || (ones == 1 && zeros == 0) || (zeros == 1 && ones == 0) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1632A(r *bufio.Reader) *CF1632A {
	return &CF1632A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1632A(bufio.NewReader(os.Stdin)).Run()
}
