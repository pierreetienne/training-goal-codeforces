package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1739A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1739A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1739A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1739A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1739A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1739A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1739A
Date: 9/29/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1739/A
Problem: CF1739A
**/
func (in *CF1739A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt(), in.NextInt()
		fmt.Println(int(math.Ceil(float64(n)/2.)), int(math.Ceil(float64(m)/2.)))
	}
}

func NewCF1739A(r *bufio.Reader) *CF1739A {
	return &CF1739A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1739A(bufio.NewReader(os.Stdin)).Run()
}
