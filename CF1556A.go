package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1556A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1556A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1556A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1556A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1556A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1556A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1556A
Date: 14/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1556/A
Problem: CF1556A
**/
func (in *CF1556A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, b := in.NextInt(), in.NextInt()

		max := int(math.Max(float64(a), float64(b)))
		min := int(math.Min(float64(a), float64(b)))

		if (max-min)%2 == 0 {
			if max == min {
				if max == 0 {
					fmt.Println(0)
				} else {
					fmt.Println(1)
				}
			} else {
				fmt.Println(2)
			}

		} else {
			fmt.Println(-1)
		}
	}
}

func NewCF1556A(r *bufio.Reader) *CF1556A {
	return &CF1556A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1556A(bufio.NewReader(os.Stdin)).Run()
}
