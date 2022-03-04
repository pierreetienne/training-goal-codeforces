package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1065A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1065A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1065A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1065A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1065A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1065A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1065A
Date: 2/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1065/A
Problem: CF1065A
**/
func (in *CF1065A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		s, a, b, c := in.NextInt64(), in.NextInt64(), in.NextInt64(), in.NextInt64()
		bars := int64(math.Floor(float64(s / c)))
		x := int64(math.Floor(float64(bars / a)))
		ans := bars + (x * b)
		fmt.Println(ans)

	}
}

func NewCF1065A(r *bufio.Reader) *CF1065A {
	return &CF1065A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1065A(bufio.NewReader(os.Stdin)).Run()
}
