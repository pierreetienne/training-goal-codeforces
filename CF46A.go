package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF46A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF46A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF46A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF46A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF46A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF46A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF46A
Date: 10/01/23
User: pepradere
URL: https://codeforces.com/problemset/problem/46/A
Problem: CF46A
**/
func (in *CF46A) Run() {
	n := in.NextInt()
	count := 0
	init := 0
	mas := 1
	for count < n-1 {
		init = (init + mas) % n
		fmt.Print(init+1, " ")
		mas++
		count++
	}
	fmt.Println()
}

func NewCF46A(r *bufio.Reader) *CF46A {
	return &CF46A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF46A(bufio.NewReader(os.Stdin)).Run()
}
