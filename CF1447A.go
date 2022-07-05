package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1447A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1447A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1447A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1447A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1447A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1447A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1447A
Date: 3/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1447/A
Problem: CF1447A
**/
func (in *CF1447A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		fmt.Println(n)
		var str strings.Builder
		for i:=0;i<n;i++{
			str.WriteString(fmt.Sprintf("%d ", i+1))
		}
		fmt.Println(str.String())
	}
}

func NewCF1447A(r *bufio.Reader) *CF1447A {
	return &CF1447A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1447A(bufio.NewReader(os.Stdin)).Run()
}
