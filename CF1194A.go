package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1194A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1194A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1194A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1194A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1194A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1194A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1194A
Date: 21/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1194/A
Problem: CF1194A
**/
func (in *CF1194A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		_, n := in.NextInt(), in.NextInt()
		fmt.Println(n * 2)
	}
}

func NewCF1194A(r *bufio.Reader) *CF1194A {
	return &CF1194A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1194A(bufio.NewReader(os.Stdin)).Run()
}
