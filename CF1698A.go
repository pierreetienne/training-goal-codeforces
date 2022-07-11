package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1698A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1698A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1698A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1698A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1698A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1698A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1698A
Date: 9/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1698/A
Problem: CF1698A
**/
func (in *CF1698A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		ans := 0
		len := in.NextInt()
		for i := 0; i < len; i++ {
			 ans = in.NextInt()
		}
		fmt.Println(ans)
	}
}

func NewCF1698A(r *bufio.Reader) *CF1698A {
	return &CF1698A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1698A(bufio.NewReader(os.Stdin)).Run()
}
