package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1771A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1771A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1771A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1771A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1771A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1771A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1771A
Date: 6/02/23
User: pepradere
URL: https://codeforces.com/problemset/problem/1771/A
Problem: CF1771A
**/
func (in *CF1771A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		min := 999999999
		cantMin := 0
		max := -100000
		cantMax := 0
		for i := 0; i < n; i++ {
			val := in.NextInt()
			if val < min {
				min = val
				cantMin = 0
			}
			if val == min {
				cantMin++
			}

			if val > max {
				max = val
				cantMax = 0
			}
			if val == max {
				cantMax++
			}

		}
		cont := 2 * (cantMin * cantMax)

		if max == min {
			cont = n * (n - 1)
		}

		fmt.Println(cont)
	}
}

func NewCF1771A(r *bufio.Reader) *CF1771A {
	return &CF1771A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1771A(bufio.NewReader(os.Stdin)).Run()
}
