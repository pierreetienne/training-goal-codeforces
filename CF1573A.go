package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1573A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1573A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1573A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1573A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1573A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1573A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1573A
Date: 30/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1573/A
Problem: CF1573A
**/
func (in *CF1573A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		ans := 0
		for i := 0; i < n; i++ {
			if str[i]-'0' > 0 {
				ans += int(str[i] - '0')
				if i+1 < n {
					ans++
				}
			}
		}
		fmt.Println(ans)
	}
}

func NewCF1573A(r *bufio.Reader) *CF1573A {
	return &CF1573A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1573A(bufio.NewReader(os.Stdin)).Run()
}
