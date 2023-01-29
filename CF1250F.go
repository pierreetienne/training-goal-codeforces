package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1250F struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1250F) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1250F) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1250F) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1250F) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1250F) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1250F
Date: 1/20/2023
User: wotan
URL: https://codeforces.com/problemset/problem/1250/F
Problem: CF1250F
**/
func (in *CF1250F) Run() {
	n := in.NextInt()

	min := n*n + 1000
	to := int(math.Sqrt(float64(n)))
	for i := 1; i <= to; i++ {
		v := n / i
		if v*i == n {
			min = int(math.Min(float64(2*v+2*i), float64(min)))
		}
	}
	fmt.Println(min)
}

func NewCF1250F(r *bufio.Reader) *CF1250F {
	return &CF1250F{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1250F(bufio.NewReader(os.Stdin)).Run()
}
