package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF897A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF897A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF897A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF897A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF897A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF897A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF897A
Date: 3/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/897/A
Problem: CF897A
**/
func (in *CF897A) Run() {
	n, m := in.NextInt(), in.NextInt()
	str := in.NextString()
	runes := make([]uint8, n)
	for i := 0; i < n; i++ {
		runes[i] = str[i]
	}
	for i := 0; i < m; i++ {
		l, r := in.NextInt()-1, in.NextInt()
		c1, c2 := in.NextString()[0], in.NextString()[0]
		for j := l; j < r; j++ {
			if runes[j] == c1 {
				runes[j] = c2
			}
		}
	}
	var sol strings.Builder
	for i:=0;i<n;i++{
		sol.WriteString(string(runes[i]))
	}
	sol.WriteString("\n")
	fmt.Print(sol.String())

}

func NewCF897A(r *bufio.Reader) *CF897A {
	return &CF897A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF897A(bufio.NewReader(os.Stdin)).Run()
}
