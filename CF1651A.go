package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1651A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1651A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1651A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1651A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1651A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1651A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1651A
Date: 21/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1651/A
Problem: CF1651A
**/
func (in *CF1651A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		p := int(math.Pow(2, float64(n)))
		fmt.Println(p - 1)
	}
}

func NewCF1651A(r *bufio.Reader) *CF1651A {
	return &CF1651A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1651A(bufio.NewReader(os.Stdin)).Run()
}
