package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1612A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1612A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1612A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1612A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1612A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1612A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1612A
Date: 10/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1612/A
Problem: CF1612A
**/
func (in *CF1612A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		x, y := in.NextInt(), in.NextInt()
		if (x+y)%2 == 0 {
			div := (x + y) / 2
			a := int(math.Floor(float64(x) / 2.))
			b := int(math.Abs(float64(a) - float64(div)))
			fmt.Println(a, b)

		} else {
			fmt.Println("-1 -1")
		}
	}
}

func NewCF1612A(r *bufio.Reader) *CF1612A {
	return &CF1612A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1612A(bufio.NewReader(os.Stdin)).Run()
}
