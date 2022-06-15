package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1428A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1428A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1428A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1428A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1428A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1428A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1428A
Date: 13/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1428/A
Problem: CF1428A
**/
func (in *CF1428A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		x1, y1, x2, y2 := in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
		ans := 0
		if x1 == x2 {
			ans = int(math.Abs(float64(y1) - float64(y2)))
		} else if y1 == y2 {
			ans = int(math.Abs(float64(x1) - float64(x2)))
		} else {
			ans = 2 + int(math.Abs(float64(x1)-float64(x2))) + int(math.Abs(float64(y1)-float64(y2)))
		}
		fmt.Println(ans)

	}
}

func NewCF1428A(r *bufio.Reader) *CF1428A {
	return &CF1428A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1428A(bufio.NewReader(os.Stdin)).Run()
}
