package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1519A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1519A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1519A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1519A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1519A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1519A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1519A
Date: 19/10/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1519/A
Problem: CF1519A
**/
func (in *CF1519A) Run() {
	t := in.NextInt()

	for ; t > 0; t-- {
		r, b, d := in.NextInt(), in.NextInt(), in.NextInt()
		min := Min(r, b)
		max := Max(r, b)
		pack := int(math.Ceil(float64(max) /float64( min)))
		if max - min <= d || int(math.Abs(float64(pack-1))) <= d  {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}

}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NewCF1519A(r *bufio.Reader) *CF1519A {
	return &CF1519A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1519A(bufio.NewReader(os.Stdin)).Run()
}
