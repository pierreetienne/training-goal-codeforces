package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1729B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1729B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1729B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1729B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1729B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1729B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1729B
Date: 9/13/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1729/B
Problem: CF1729B
**/
func (in *CF1729B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		sol := ""
		for i := n - 1; i >= 0; {
			if str[i] == '0' && i-2 >= 0 {
				val := int(str[i-1]-'0') + int(str[i-2]-'0')*10
				sol = string(rune(val+'a'-1)) + sol
				i -= 3
			} else {
				val := str[i]
				sol = string(rune(val-'0'+'a'-1)) + sol
				i -= 1
			}
		}
		fmt.Println(sol)
	}
}

func NewCF1729B(r *bufio.Reader) *CF1729B {
	return &CF1729B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1729B(bufio.NewReader(os.Stdin)).Run()
}
