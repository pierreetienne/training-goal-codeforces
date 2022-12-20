package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1774A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1774A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1774A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1774A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1774A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1774A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1774A
Date: 12/18/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1774/A
Problem: CF1774A
**/
func (in *CF1774A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		ans := int(str[0] - '0')
		var sol strings.Builder

		for i := 1; i < n; i++ {
			if ans == 1 {
				sol.WriteString("-")
				ans -= int(str[i] - '0')
			} else if ans == -1 {
				sol.WriteString("+")
				ans += int(str[i] - '0')
			} else {
				sol.WriteString("+")
				ans += int(str[i] - '0')
			}
		}
		fmt.Println(sol.String())
	}
}

func NewCF1774A(r *bufio.Reader) *CF1774A {
	return &CF1774A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1774A(bufio.NewReader(os.Stdin)).Run()
}
