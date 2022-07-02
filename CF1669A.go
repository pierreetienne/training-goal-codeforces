package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1669A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1669A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1669A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1669A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1669A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1669A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1669A
Date: 23/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1669/A
Problem: CF1669A
**/
func (in *CF1669A) Run() {
	var ans strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		r := in.NextInt()
		if r <= 1399 {
			ans.WriteString("Division 4\n")
		} else if r <= 1599 {
			ans.WriteString("Division 3\n")
		} else if r <= 1899{
			ans.WriteString("Division 2\n")
		} else {
			ans.WriteString("Division 1\n")
		}

	}
	fmt.Print(ans.String())
}

func NewCF1669A(r *bufio.Reader) *CF1669A {
	return &CF1669A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1669A(bufio.NewReader(os.Stdin)).Run()
}
