package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF859A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF859A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF859A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF859A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF859A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF859A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF859A
Date: 10/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/859/A
Problem: CF859A
**/
func (in *CF859A) Run() {
	n := in.NextInt()
	max := 0
	for i := 0; i < n; i++ {
		val := in.NextInt()
		if val > max {
			max = val
		}
	}

	if max > 25 {
		fmt.Println(max - 25)
	} else {
		fmt.Println(0)
	}
}

func NewCF859A(r *bufio.Reader) *CF859A {
	return &CF859A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF859A(bufio.NewReader(os.Stdin)).Run()
}
