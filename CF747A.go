package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF747A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF747A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF747A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF747A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF747A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF747A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF747A
Date: 4/09/22
User: pepradere
URL: https://codeforces.com/problemset/problem/747/A
Problem: CF747A
**/
func (in *CF747A) Run() {
	n := in.NextInt()

	r := int(math.Sqrt(float64(n)))
	v := n / r
	for v*r != n && v > 0 {
		r--
		v = n / r
	}
	fmt.Println(r, v)
}

func NewCF747A(r *bufio.Reader) *CF747A {
	return &CF747A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF747A(bufio.NewReader(os.Stdin)).Run()
}
