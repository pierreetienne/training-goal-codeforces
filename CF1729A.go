package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1729A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1729A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1729A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1729A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1729A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1729A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1729A
Date: 12/09/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1729/A
Problem: CF1729A
**/
func (in *CF1729A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, b, c := in.NextInt(), in.NextInt(), in.NextInt()
		diffA := a

		diffB := int(math.Abs(float64(b)-float64(c))) + c
		if diffA == diffB {
			fmt.Println(3)
		} else if diffA < diffB {
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}
	}
}

func NewCF1729A(r *bufio.Reader) *CF1729A {
	return &CF1729A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1729A(bufio.NewReader(os.Stdin)).Run()
}
