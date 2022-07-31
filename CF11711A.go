package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF11711A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF11711A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF11711A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF11711A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF11711A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF11711A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF11711A
Date: 29/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1711/A
Problem: CF11711A
**/
func (in *CF11711A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		ans := strings.Builder{}
		ans.WriteString(fmt.Sprintf("%d", n))
		for i := 1; i < n; i++ {
			ans.WriteString(fmt.Sprintf(" %d", i))
		}
		fmt.Println(ans.String())
	}
}

func NewCF11711A(r *bufio.Reader) *CF11711A {
	return &CF11711A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF11711A(bufio.NewReader(os.Stdin)).Run()
}
