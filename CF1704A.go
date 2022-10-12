package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1704A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1704A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1704A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1704A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1704A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1704A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1704A
Date: 10/10/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1704/A
Problem: CF1704A
**/
func (in *CF1704A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt(), in.NextInt()
		a := in.NextString()
		b := in.NextString()

		ans := true
		if m > n {
			ans = false
		} else {
			eq := 0
			for i, j := n-1, m-1; j >= 0 && i >= 0; {
				if a[i] == b[j] {
					eq++
					j--
					i--
				} else if eq == m-1 {
					i--
				} else {
					break
				}
			}
			if eq < m {
				ans = false
			}
		}

		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func NewCF1704A(r *bufio.Reader) *CF1704A {
	return &CF1704A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1704A(bufio.NewReader(os.Stdin)).Run()
}
