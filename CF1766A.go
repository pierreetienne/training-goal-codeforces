package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1766A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1766A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1766A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1766A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1766A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1766A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1766A
Date: 20/12/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1766/A
Problem: CF1766A
**/
func (in *CF1766A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		ans := int(str[0]-'0') + (9 * (len(str) - 1))
		fmt.Println(ans)
	}
}

func NewCF1766A(r *bufio.Reader) *CF1766A {
	return &CF1766A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1766A(bufio.NewReader(os.Stdin)).Run()
}
