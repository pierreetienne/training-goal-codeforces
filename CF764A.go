package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF764A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF764A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF764A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF764A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF764A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF764A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF764A
Date: 3/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/764/A
Problem: CF764A
**/
func (in *CF764A) Run() {
	n, m, z := in.NextInt(), in.NextInt(), in.NextInt()
	kill := 0
	for i := 1; i <= z; i++ {
		if i%n == 0 && i%m == 0 {
			kill++
		}
	}
	fmt.Println(kill)
}

func NewCF764A(r *bufio.Reader) *CF764A {
	return &CF764A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF764A(bufio.NewReader(os.Stdin)).Run()
}
