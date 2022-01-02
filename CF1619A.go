package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1619A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1619A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1619A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1619A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1619A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1619A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1619A
Date: 1/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1619/A
Problem: CF1619A
**/
func (in *CF1619A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		ans := "NO"
		if len(str)%2 == 0 {
			if str[0:len(str)/2] == str[len(str)/2:]{
				ans = "YES"
			}
		}
		fmt.Println(ans)
	}
}

func NewCF1619A(r *bufio.Reader) *CF1619A {
	return &CF1619A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1619A(bufio.NewReader(os.Stdin)).Run()
}
