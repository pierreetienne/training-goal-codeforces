package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF952A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF952A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF952A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF952A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF952A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF952A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF952A
Date: 13/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/952/A
Problem: CF952A
**/
func (in *CF952A) Run() {
	fmt.Println(in.NextInt()%2)
}

func NewCF952A(r *bufio.Reader) *CF952A {
	return &CF952A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF952A(bufio.NewReader(os.Stdin)).Run()
}
