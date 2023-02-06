package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1475B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1475B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1475B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1475B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1475B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1475B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1475B
Date: 2/2/2023
User: wotan
URL: https://codeforces.com/problemset/problem/1475/B
Problem: CF1475B
**/
func (in *CF1475B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		a := int(math.Ceil(float64(n)/2020.)) + 1
		f := false
		for a >= 0 {
			val := n - (a * 2020)
			if val%2021 == 0 && val != 1 && val >= 0 {
				f = true
				break
			} else {
				a--
			}
		}

		if f {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func NewCF1475B(r *bufio.Reader) *CF1475B {
	return &CF1475B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1475B(bufio.NewReader(os.Stdin)).Run()
}
