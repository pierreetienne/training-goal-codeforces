package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1621A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1621A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1621A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1621A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1621A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1621A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1621A
Date: 19/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1621/A
Problem: CF1621A
**/
func (in *CF1621A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()

		if (k*2)-1 > n {
			fmt.Println(-1)
		} else {
			r := 0
			var str strings.Builder
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if i == r && j == r && k > 0 {
						str.WriteString("R")
						r += 2
						k--
					} else {
						str.WriteString(".")
					}
				}
				str.WriteString("\n")
			}
			fmt.Print(str.String())
		}
	}
}

func NewCF1621A(r *bufio.Reader) *CF1621A {
	return &CF1621A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1621A(bufio.NewReader(os.Stdin)).Run()
}
